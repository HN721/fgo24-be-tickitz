package models

import (
	"context"
	"time"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

type Movies struct {
	Title       string    `db:"title"`
	Synopsis    string    `db:"synopsis"`
	Background  string    `db:"background"`
	Poster      string    `db:"poster"`
	ReleaseDate time.Time `db:"release_date"`
	Duration    int       `db:"duration"`
	Price       int       `db:"price"`
}

func GetAllMovies() ([]Movies, error) {
	conn, err := utils.DBConnect()
	if err != nil {

	}
	query := `SELECT title, synopsis, background, poster, release_date, duration, price FROM movies`
	result, err := conn.Query(context.Background(), query)
	data, err := pgx.CollectRows[Movies](result, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}
	return data, nil

}
func InsertMovies(movie Movies) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	query := `INSERT INTO movies(title, synopsis, background, poster, release_date, duration, price)VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err = conn.Exec(context.Background(), query, movie.Title, movie.Synopsis, movie.Background, movie.Poster, movie.ReleaseDate, movie.Duration, movie.Price)
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
