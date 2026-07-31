package rest

import (
	http "net/http"
)

const prefix = "/players"

func RegisterRoutes(mux *http.ServeMux) {
	createPlayerContainer := NewContainer()

	mux.HandleFunc("POST "+prefix+"/create", createPlayerContainer.CreateHandler.CreatePlayerHttpHandler)
	mux.HandleFunc(prefix+"/get", GetPlayerByIdHttpHandler)
}
