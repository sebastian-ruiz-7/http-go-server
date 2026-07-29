package domain

import (
	"strings"
)

type Player struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ValidatePlayerData(player Player) error {
	return nil
}

func ParsePlayerData(player *Player) {
	player.Name = strings.TrimSpace(player.Name)
	// player.Age=number.

}
