package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hahnap4/social-media/http-server/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

type errorBody struct {
	Error string `json:"error"`
}

func main() {
	m := http.NewServeMux()

	c := database.NewClient("db.json")
	dbErr := c.EnsureDB()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	apiCfg := apiConfig{
		dbClient: c,
	}

	m.HandleFunc("/users", apiCfg.endpointUsersHandler)
	m.HandleFunc("/users/", apiCfg.endpointUsersHandler)
	m.HandleFunc("/posts", apiCfg.endpointPostsHandler)
	m.HandleFunc("/posts/", apiCfg.endpointPostsHandler)

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
