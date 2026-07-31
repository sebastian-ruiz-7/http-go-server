package domain

type Player struct {
	ID   string
	Name string
	Age  int
}

func CreatePlayer(player Player) error {
	err := validateDataForCreatePlayer(player)

	return err
}
