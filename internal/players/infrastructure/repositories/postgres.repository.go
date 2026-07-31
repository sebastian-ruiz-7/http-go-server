package repositories

import (
	"fmt"

	"github.com/sebastian-ruiz-7/http-go-server/internal/players/domain"
)

// type PostgresRepository interface {
// 	Save(player domain.Player) error
// }

type PostgresRepository struct {
}

func (pr PostgresRepository) Save(player domain.Player) error {
	fmt.Println("saving repository", player)
	return nil
}
