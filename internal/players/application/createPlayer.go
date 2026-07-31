package application

import (
	"github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
	userDomain "github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
)

type CreatePlayerService struct {
	ids  userDomain.IDGenerator
	repo userDomain.PlayerRepository
}

func NewCreatePlayerService(ids userDomain.IDGenerator, repo userDomain.PlayerRepository) *CreatePlayerService {
	return &CreatePlayerService{ids: ids, repo: repo}
}

type CreatePlayerInput struct {
	Name string
	Age  int
}

type CreatePlayerUseCase interface {
	Execute(inputData CreatePlayerInput) error
}

func (us *CreatePlayerService) Execute(inputData CreatePlayerInput) error {
	playerData := domain.Player{ID: us.ids.NewID(), Name: inputData.Name, Age: inputData.Age}

	err := domain.CreatePlayer(playerData)

	if err != nil {
		return err
	}

	err = us.repo.Save(playerData)

	if err != nil {
		return err
	}

	return nil
}
