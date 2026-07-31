package domain

type PlayerRepository interface {
	Save(player Player) error
}
