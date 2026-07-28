package rest

import (
	http "net/http"
)

const prefix = "/players"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST "+prefix+"/create", CreatePlayerHttpHandler)
	mux.HandleFunc(prefix+"/get", GetPlayerByIdHttpHandler)
}
