package main

import (
	"net/http"
	"time"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}

func main() {

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", testHandler)

	addr := "localhost:8080"
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	http.ListenAndServe(srv.Addr, serveMux)
}
