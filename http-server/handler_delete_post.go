package main

import (
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	uuid := strings.TrimPrefix(r.URL.Path, "/posts/")
	if uuid == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("handlerDeletePost: no id provided"))
		return
	}

	err := apiCfg.dbClient.DeletePost(uuid)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
