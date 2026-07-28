package rest

import (
	"fmt"
	"io"
	"net/http"

	createPlayerUserCase "github.com/sebastian-ruiz-7/http-go-server/internal/players/application"
)

func CreatePlayerHttpHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("creating a player")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("error en io", err)
	}

	createPlayerUserCase.CreatePlayerUseCase(body)
	w.Write(message)
}
