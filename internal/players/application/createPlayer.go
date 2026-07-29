package application

import (
	"encoding/json"
	"fmt"
	"strings"

	playersDomain "github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
)

func CreatePlayerUseCase(body []byte) {
	//trying deserializing json

	var player playersDomain.Player

	err := json.Unmarshal(body, &player)

	// fmt.Println("err", err)

	if err != nil {
		fmt.Println("error pa")
	}

	fmt.Println("longitud", len(player.Name))
	player.Name = strings.TrimSpace(player.Name)

	fmt.Println("edad", player.Age)
}
