package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func DBConnect() (*pgxpool.Conn, error) {
	godotenv.Load()
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		"postgres",
		"721",
		"db",
		"5432",
		"movxtar",
	)
	pool, err := pgxpool.New(
		context.Background(), connection,
	)
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		return conn, err
	}
	return conn, err
}
