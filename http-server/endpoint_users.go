package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:
		apiCfg.handlerCreateUser(w, r)
	case http.MethodPut:

	case http.MethodDelete:

	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}
