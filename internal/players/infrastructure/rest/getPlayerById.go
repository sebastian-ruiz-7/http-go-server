package rest

import "net/http"

func GetPlayerByIdHttpHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("getting user")

	w.Write(message)
}
