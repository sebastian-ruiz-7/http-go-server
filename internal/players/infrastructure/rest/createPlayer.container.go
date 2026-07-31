package rest

import (
	"github.com/sebastian-ruiz-7/http-go-server/internal/players/application"
	"github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
	IDS "github.com/sebastian-ruiz-7/http-go-server/internal/players/infrastructure"
	Reposotories "github.com/sebastian-ruiz-7/http-go-server/internal/players/infrastructure/repositories"
)

type CreatePlayerContainer struct {
	IDGenerator   domain.IDGenerator
	CreatePlayer  application.CreatePlayerUseCase
	CreateHandler CreatePlayerHandler
	RepoPlayer    domain.PlayerRepository
}

func NewContainer() *CreatePlayerContainer {
	ids := IDS.UUIDGenerator{}
	postgresRepo := Reposotories.PostgresRepository{}

	createPlayer := application.NewCreatePlayerService(ids, postgresRepo)

	createPlayerHandler := NewCreatePlayerHandler(createPlayer)

	return &CreatePlayerContainer{
		IDGenerator:   ids,
		CreatePlayer:  createPlayer,
		CreateHandler: *createPlayerHandler,
	}
}
