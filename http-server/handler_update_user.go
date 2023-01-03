package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	email := strings.TrimPrefix(r.URL.Path, "/users/")
	if email == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("handlerUpdateUser: no email provided"))
		return
	}

	data, err := ioutil.ReadAll(r.Body)

	params := parameters{}

	json.Unmarshal(data, &params)

	fmt.Println(params)

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

	updatedUser, err := apiCfg.dbClient.UpdateUser(
		email,
		params.Password,
		params.Name,
		params.Age,
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedUser)
}
