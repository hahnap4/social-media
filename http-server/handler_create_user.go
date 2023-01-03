package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
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

	newUser, err := apiCfg.dbClient.CreateUser(
		params.Email,
		params.Password,
		params.Name,
		params.Age,
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, newUser)
}
