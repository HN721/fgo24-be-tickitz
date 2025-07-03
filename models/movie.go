package models

import (
	"context"
	"fmt"
	"time"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

type Movies struct {
	Id          int       `json:"id,omitempty" db:"id"`
	Title       string    `json:"title" db:"title"`
	Synopsis    string    `json:"synopsis" db:"synopsis"`
	Background  string    ` json:"background" db:"background"`
	Poster      string    ` json:"poster" db:"poster"`
	ReleaseDate time.Time ` json:"releaseDate" db:"release_date"`
	Duration    int       `json:"duration" db:"duration"`
	Price       int       `json:"price" db:"price"`
}
type Genres struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type Actor struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type Directors struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func GetUpcomingMovies() ([]Movies, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `
	SELECT id, title, synopsis, background, poster, release_date, duration, price
	FROM movies
	WHERE release_date > CURRENT_DATE
	`)
	if err != nil {
		return nil, err
	}
	results, err := pgx.CollectRows[Movies](rows, pgx.RowToStructByName)
	return results, err

}
func GetAllMovies() ([]Movies, error) {
	conn, err := utils.DBConnect()
	if err != nil {

	}
	query := `SELECT id,title, synopsis, background, poster, release_date, duration, price FROM movies`
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
	if movie.Title == "" {
		return fmt.Errorf("Judul Tidak Boleh Kosong")
	}
	query := `
		INSERT INTO movies (title, synopsis, background, poster, release_date, duration, price)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err = conn.Exec(context.Background(), query,
		movie.Title,
		movie.Synopsis,
		movie.Background,
		movie.Poster,
		movie.ReleaseDate,
		movie.Duration,
		movie.Price,
	)
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
func UpdateMovies(movie Movies, movieId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	query := `
	UPDATE movies 
	SET title = $1, synopsis = $2, background = $3, poster = $4, release_date = $5, duration = $6, price = $7
	WHERE id = $8`

	result, err := conn.Exec(context.Background(), query,
		movie.Title,
		movie.Synopsis,
		movie.Background,
		movie.Poster,
		movie.ReleaseDate,
		movie.Duration,
		movie.Price,
		movieId,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("movie with id %d not found", movieId)
	}

	return nil
}

func DeleteMovies(movieId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	query := `DELETE FROM movies WHERE id =$1`
	result, err := conn.Exec(context.Background(), query, movieId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("tidak ada movie dengan id %d", movieId)
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return nil

}

// genres
func GenreMovies() ([]Genres, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	query := `SELECT id,name FROM genres`
	result, err := conn.Query(context.Background(), query)
	data, err := pgx.CollectRows[Genres](result, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return data, nil

}
func CreateGenre(genre Genres) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	if genre.Name == "" {
		return fmt.Errorf("Nama Tidak Boleh Kosong")
	}
	query := `INSERT INTO genres(name)VALUES($1)`
	_, err = conn.Exec(context.Background(), query, genre.Name)
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
func UpdateGenre(genre Genres, genreId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	if genre.Name == "" {
		return fmt.Errorf("Nama Tidak Boleh Kosong")
	}
	query := `
	UPDATE genres 
	SET name = $1
	WHERE id = $2`
	results, err := conn.Exec(context.Background(), query, genre.Name, genreId)
	if err != nil {
		return err
	}

	if results.RowsAffected() == 0 {
		return fmt.Errorf("movie with id %d not found", genreId)
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
func DeleteGenre(genreId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	query := `DELETE FROM genres WHERE id =$1`
	result, err := conn.Exec(context.Background(), query, genreId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("tidak ada genre dengan id %d", genreId)
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}

// Actors
func ActorMovies() ([]Actor, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	query := `SELECT id,fullname FROM actors`
	result, err := conn.Query(context.Background(), query)
	data, err := pgx.CollectRows[Actor](result, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return data, nil

}
func CreateActor(actor Actor) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	if actor.Name == "" {
		return fmt.Errorf("Nama Tidak Boleh Kosong")
	}
	query := `INSERT INTO actors(fullname)VALUES($1)`
	_, err = conn.Exec(context.Background(), query, actor.Name)
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
func UpdateActor(actor Actor, actorId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	if actor.Name == "" {
		return fmt.Errorf("Nama Tidak Boleh Kosong")
	}
	query := `
	UPDATE actors 
	SET fullname = $1
	WHERE id = $2`
	results, err := conn.Exec(context.Background(), query, actor.Name, actorId)
	if err != nil {
		return err
	}

	if results.RowsAffected() == 0 {
		return fmt.Errorf("movie with id %d not found", actorId)
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
func DeleteActor(actorId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	query := `DELETE FROM actors WHERE id =$1`
	result, err := conn.Exec(context.Background(), query, actorId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("tidak ada actor dengan id %d", actorId)
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}

// Directors

func DirectorsMovie() ([]Directors, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	query := `SELECT id,fullname FROM directors`
	result, err := conn.Query(context.Background(), query)
	data, err := pgx.CollectRows[Directors](result, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return data, nil

}
func CreateDirector(director Directors) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	if director.Name == "" {
		return fmt.Errorf("Nama Tidak Boleh Kosong")
	}
	query := `INSERT INTO directors(fullname)VALUES($1)`
	_, err = conn.Exec(context.Background(), query, director.Name)
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
func UpdateDirector(director Directors, directorId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	if director.Name == "" {
		return fmt.Errorf("Nama Tidak Boleh Kosong")
	}
	query := `
	UPDATE directors 
	SET fullname = $1
	WHERE id = $2`
	results, err := conn.Exec(context.Background(), query, director.Name, directorId)
	if err != nil {
		return err
	}

	if results.RowsAffected() == 0 {
		return fmt.Errorf("director with id %d not found", directorId)
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
func DeleteDirector(directorId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	query := `DELETE FROM directors WHERE id =$1`
	result, err := conn.Exec(context.Background(), query, directorId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("tidak ada director dengan id %d", directorId)
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
