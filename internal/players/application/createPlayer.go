package application

import (
	"fmt"

	userDomain "github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
)

// interfaz
type CreatePlayerUseCase interface {
	CreatePlayer(player userDomain.Player) error
}

// Puerto de salida
type IDGenerator interface {
	NewId() string
}

// Inyección de dependencia del IDs
type CreatePlayerService struct {
	ids IDGenerator
}

func NewCreatePlayerService(ids IDGenerator) *CreatePlayerService {
	return &CreatePlayerService{ids: ids}
}

func CreatePlayer(player userDomain.Player) {
	fmt.Println("player", player)
}
