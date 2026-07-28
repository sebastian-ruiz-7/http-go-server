package rest

import (
	"net/http"
)

func CreateTeamHttpHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("creating a team")

	w.Write(message)
}
