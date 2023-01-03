package main

import (
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	if email == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("handlerGetUser: no email provided"))
		return
	}

	user, err := apiCfg.dbClient.GetUser(email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
	}

	respondWithJSON(w, http.StatusOK, user)
}
