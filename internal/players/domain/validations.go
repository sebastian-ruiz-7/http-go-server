package domain

import "errors"

func ValidateDataForCreatePlayer(player Player) error {
	if len(player.Name) == 0 {
		return errors.New("Name can't be empty")
	}

	if player.Age == 0 {
		return errors.New("Age can't be 0")
	}

	if player.Age < 0 {
		return errors.New("Age can't be negative")
	}
	return nil
}
