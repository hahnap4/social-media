package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hahnap4/social-media/http-server/internal/database"
)

func main() {

	m := http.NewServeMux()

	m.HandleFunc("/", testHandler)
	m.HandleFunc("/err", testErrHandler)

	addr := "localhost:8080"
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// blocks until server has unrecoverable error
	fmt.Println("server started on ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)
			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})

			w.Write(response)
			return
		}

		w.WriteHeader(code)
		w.Write(response)
	}
}

type errorBody struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("don't call respondWithError with nil error")
		return
	}

	log.Println(err)

	errorMessage := errorBody{
		Error: err.Error(),
	}

	respondWithJSON(w, code, errorMessage)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("server error"))

}
