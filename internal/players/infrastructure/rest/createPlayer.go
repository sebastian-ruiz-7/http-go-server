package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sebastian-ruiz-7/http-go-server/internal/players/application"
)

type createPlayer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createPlayerHttpAdapter(body io.Reader) (application.CreatePlayerInput, error) {
	var playerJson createPlayer

	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&playerJson); err != nil {
		var ute *json.UnmarshalTypeError
		var se *json.SyntaxError

		switch {
		case errors.As(err, &ute):
			return application.CreatePlayerInput{}, fmt.Errorf(
				"invalid type for field %q: got %q, expected %s",
				ute.Field, ute.Value, ute.Type.String(),
			)
		case errors.As(err, &se):
			return application.CreatePlayerInput{}, fmt.Errorf("invalid JSON syntax at byte %d", se.Offset)
		case strings.HasPrefix(err.Error(), "json: unkown field"):
			field := strings.TrimPrefix(err.Error(), "json: unknown field ")
			field = strings.Trim(field, "\"")
			return application.CreatePlayerInput{}, fmt.Errorf("unknown field: %s", field)
		default:
			if strings.Contains(err.Error(), "unknown field") {
				return application.CreatePlayerInput{}, fmt.Errorf("%v", err)
			}
			return application.CreatePlayerInput{}, fmt.Errorf("invalid JSON: %v", err)
		}
	}

	playerData := application.CreatePlayerInput{Name: playerJson.Name, Age: playerJson.Age}
	return playerData, nil
}

type CreatePlayerHandler struct {
	useCase application.CreatePlayerUseCase
}

func NewCreatePlayerHandler(useCase application.CreatePlayerUseCase) *CreatePlayerHandler {
	return &CreatePlayerHandler{useCase: useCase}
}

func (handler CreatePlayerHandler) CreatePlayerHttpHandler(w http.ResponseWriter, r *http.Request) {
	data, err := createPlayerHttpAdapter(r.Body)

	if err != nil {
		responseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.useCase.Execute(data)

	if err != nil {
		responseWriter(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func responseWriter(w http.ResponseWriter, message string, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")

	msg := map[string]string{"message": message}
	jsonData, _ := json.Marshal(msg)

	w.WriteHeader(httpStatus)
	w.Write(jsonData)
}
