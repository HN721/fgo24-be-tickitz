package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnect() (*pgxpool.Pool, error) {
	host := os.Getenv("PGHOST")
	if host == "" {
		host = "db" // docker service name
	}

	port := os.Getenv("PGPORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("PGUSER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")
	if dbname == "" {
		dbname = "postgres"
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	// test koneksi
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("✅ Connected to PostgreSQL:", host)
	return pool, nil
}
