package application

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"`
}

func CreatePlayerUseCase(body []byte) {
	//trying deserializing json

	var player Player

	err := json.Unmarshal(body, &player)

	if err != nil {
		fmt.Println("error pa")
	}

	fmt.Println("body service", player.Name)
}
