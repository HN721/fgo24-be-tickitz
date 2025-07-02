package models

import (
	"context"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

type movies struct {
	title        string
	synopsis     string
	background   string
	poster       string
	release_date string
	duration     int
	price        int
}

func GetAllMovies() []movies {
	conn, err := utils.DBConnect()
	if err != nil {

	}
	query := `SELECT id,title,synopsis,background,poster,release_date,duration,price FROM movies`
	result, err := conn.Query(context.Background(), query)
	data, err := pgx.CollectRows[movies](result, pgx.RowToStructByName)
	return data

}
