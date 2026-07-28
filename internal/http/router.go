package http

import (
	http "net/http"

	userRoute "github.com/sebastian-ruiz-7/http-go-server/internal/user/insfrastructure/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	userRoute.RegisterRoutes(mux)

	return mux
}
