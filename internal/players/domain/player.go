package domain

import "fmt"

type Player struct {
	ID   string
	Name string
	Age  int
}

func CreatePlayer(player Player) {
	fmt.Println("ID", player.ID)
	fmt.Println("Name", player.Name)
	fmt.Println("Age", player.Age)
}
