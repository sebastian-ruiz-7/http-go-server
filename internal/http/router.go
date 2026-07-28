package http

import (
	"net/http"

	playersRoutes "github.com/sebastian-ruiz-7/http-go-server/internal/players/infrastructure/rest"
	teamsRoutes "github.com/sebastian-ruiz-7/http-go-server/internal/teams/infrastructure/rest"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	playersRoutes.RegisterRoutes(mux)
	teamsRoutes.RegisterRoutes(mux)

	return mux
}
