package containers

import (
	"fmt"

	"github.com/sebastian-ruiz-7/http-go-server/internal/players/application"
	"github.com/sebastian-ruiz-7/http-go-server/internal/players/infrastructure"
	"github.com/sebastian-ruiz-7/http-go-server/internal/players/infrastructure/rest"
)

type CreatePlayerContainer struct {
	IDGenerator   application.IDGenerator
	CreatePlayer  application.CreatePlayerUseCase
	CreateHandler rest.CreatePlayerHandler
}

func NewContainer() *CreatePlayerContainer {
	ids := infrastructure.UUIDGenerator{}

	createPlayer := application.NewCreatePlayerService(ids)

	createPlayerHandler := rest.NewCreatePlayerHandler(createPlayer)

	fmt.Println(createPlayerHandler)
}
