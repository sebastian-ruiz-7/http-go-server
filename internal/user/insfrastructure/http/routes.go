package http

import (
	http "net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/hello", CreateUserHttpHandler)
}
