package shared

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConnection struct {
	pool *pgxpool.Pool
}

func NewPostgresConnection() *pgxpool.Pool {
	ctx, err := context.WithTimeout(context.Background(), 2*time.Second)

	if err != nil {
		fmt.Println("Error al crear el context para la base de datos")
	}

	pgxpool.New(ctx)
}
