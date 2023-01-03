package main

import (
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	if email == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("handlerDeleteUser: no email provided"))
		return
	}

	err := apiCfg.dbClient.DeleteUser(email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
