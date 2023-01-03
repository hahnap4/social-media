package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		switch {
		case errors.Is(err, io.ErrUnexpectedEOF):
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("Request body contains badly-formed JSON"))
		case errors.Is(err, io.EOF):
			respondWithError(w, http.StatusBadRequest, fmt.Errorf("Request body must not be empty"))
		default:
			log.Print(err.Error())
		}
	}

	newPost, err := apiCfg.dbClient.CreatePost(
		params.UserEmail,
		params.Text,
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, newPost)
}
