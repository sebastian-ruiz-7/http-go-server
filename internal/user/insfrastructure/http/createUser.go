package http

import "net/http"

func CreateUserHttpHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("creating a user")

	w.Write(message)
}
