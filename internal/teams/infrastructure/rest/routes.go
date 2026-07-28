package rest

import "net/http"

const prefix = "/teams"

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc(prefix+"/create", CreateTeamHttpHandler)
}
