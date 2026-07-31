package rest

import (
	"fmt"

	"github.com/sebastian-ruiz-7/http-go-server/internal/players/application"
	"github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
	IDS "github.com/sebastian-ruiz-7/http-go-server/internal/players/infrastructure"
)

type CreatePlayerContainer struct {
	IDGenerator   domain.IDGenerator
	CreatePlayer  application.CreatePlayerUseCase
	CreateHandler CreatePlayerHandler
}

func NewContainer() *CreatePlayerContainer {
	ids := IDS.UUIDGenerator{}

	createPlayer := application.NewCreatePlayerService(ids)

	createPlayerHandler := NewCreatePlayerHandler(createPlayer)

	fmt.Println(createPlayerHandler)
	return &CreatePlayerContainer{
		IDGenerator:   ids,
		CreatePlayer:  createPlayer,
		CreateHandler: *createPlayerHandler,
	}
}
