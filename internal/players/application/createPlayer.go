package application

import (
	"fmt"

	"github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
	userDomain "github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
)

// interfaz
// type CreatePlayerUseCase interface {
// 	CreatePlayer(player userDomain.Player) error
// }

// // Puerto de salida
// type IDGenerator interface {
// 	NewId() string
// }

// // Inyección de dependencia del IDs
// type CreatePlayerService struct {
// 	ids IDGenerator
// }

// func NewCreatePlayerService(ids IDGenerator) *CreatePlayerService {
// 	return &CreatePlayerService{ids: ids}
// }

// func CreatePlayer(player userDomain.Player) {
// 	fmt.Println("player", player)
// }

type CreatePlayerService struct {
	ids userDomain.IDGenerator
}

func NewCreatePlayerService(ids userDomain.IDGenerator) *CreatePlayerService {
	return &CreatePlayerService{ids: ids}
}

type CreatePlayerInput struct {
	Name string
	Age  int
}

type CreatePlayerUseCase interface {
	Execute(inputData CreatePlayerInput) error
}

func (us *CreatePlayerService) Execute(inputData CreatePlayerInput) error {
	fmt.Println("entro")
	id := us.ids.NewID()

	playerData := domain.Player{ID: id, Name: inputData.Name, Age: inputData.Age}

	domain.CreatePlayer(playerData)
	return nil
}
