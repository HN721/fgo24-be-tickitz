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
	Genres      []string  `json:"genres" `
	Casts       []string  `json:"casts" `
	Directors   []string  `json:"directors" `
}
type MoviesReq struct {
	Title       string `form:"title" json:"title"`
	ReleaseDate string `form:"releaseDate" json:"releaseDate"`
	Duration    int    `form:"duration" json:"duration"`
	Synopsis    string `form:"description" json:"description"`
	Price       int    `form:"price" json:"price"`
	Poster      string `form:"poster"`
	Background  string `form:"background"`
	Casts       []int  `form:"casts[]"`
	Genres      []int  `form:"genres[]"`
	Directors   []int  `form:"directors[]"`
}
type Genres struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type Actor struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"fullname" db:"fullname"`
}
type Directors struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"fullname" db:"fullname"`
}

func FilterMoviesByGenre(genreName string) ([]Movies, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
	SELECT 
		m.id,
		m.title,
		m.synopsis,
		m.background,
		m.poster,
		m.release_date,
		m.duration,
		m.price,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT g.name), NULL) AS genres,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT a.fullname), NULL) AS casts,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT d.fullname), NULL) AS directors
	FROM movies m
	LEFT JOIN movie_genre mg ON m.id = mg.movie_id
	LEFT JOIN genres g ON mg.genre_id = g.id
	LEFT JOIN movie_actors ma ON m.id = ma.movie_id
	LEFT JOIN actors a ON ma.actor_id = a.id
	LEFT JOIN movie_director md ON m.id = md.movie_id
	LEFT JOIN directors d ON md.director_id = d.id
	WHERE LOWER(g.name) LIKE LOWER($1)
	GROUP BY m.id, m.title, m.synopsis, m.background, m.poster, m.release_date, m.duration, m.price
	`

	genrePattern := "%" + genreName + "%"
	rows, err := conn.Query(context.Background(), query, genrePattern)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows[Movies](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetUpcomingMovies() ([]Movies, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
	SELECT 
		m.id,
		m.title,
		m.synopsis,
		m.background,
		m.poster,
		m.release_date,
		m.duration,
		m.price,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT g.name), NULL) AS genres,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT a.fullname), NULL) AS casts,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT d.fullname), NULL) AS directors
	FROM movies m
	LEFT JOIN movie_genre mg ON m.id = mg.movie_id
	LEFT JOIN genres g ON mg.genre_id = g.id
	LEFT JOIN movie_actors ma ON m.id = ma.movie_id
	LEFT JOIN actors a ON ma.actor_id = a.id
	LEFT JOIN movie_director md ON m.id = md.movie_id
	LEFT JOIN directors d ON md.director_id = d.id
	WHERE m.release_date > CURRENT_DATE
	GROUP BY m.id, m.title, m.synopsis, m.background, m.poster, m.release_date, m.duration, m.price
	`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows[Movies](rows, pgx.RowToStructByName)
	return results, err
}
func NowShowingMovies() ([]Movies, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
	SELECT 
		m.id,
		m.title,
		m.synopsis,
		m.background,
		m.poster,
		m.release_date,
		m.duration,
		m.price,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT g.name), NULL) AS genres,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT a.fullname), NULL) AS casts,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT d.fullname), NULL) AS directors
	FROM movies m
	LEFT JOIN movie_genre mg ON m.id = mg.movie_id
	LEFT JOIN genres g ON mg.genre_id = g.id
	LEFT JOIN movie_actors ma ON m.id = ma.movie_id
	LEFT JOIN actors a ON ma.actor_id = a.id
	LEFT JOIN movie_director md ON m.id = md.movie_id
	LEFT JOIN directors d ON md.director_id = d.id
	WHERE m.release_date < CURRENT_DATE
	GROUP BY m.id, m.title, m.synopsis, m.background, m.poster, m.release_date, m.duration, m.price
	`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows[Movies](rows, pgx.RowToStructByName)
	return results, err
}

func GetAllMovies(search string, genre string, page, limit int) ([]Movies, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	offset := (page - 1) * limit

	baseQuery := `
    SELECT 
        m.id,
        m.title,
        m.synopsis,
        m.background,
        m.poster,
        m.release_date,
        m.duration,
        m.price,
        ARRAY_REMOVE(ARRAY_AGG(DISTINCT g.name), NULL) AS genres,
        ARRAY_REMOVE(ARRAY_AGG(DISTINCT a.fullname), NULL) AS casts,
        ARRAY_REMOVE(ARRAY_AGG(DISTINCT d.fullname), NULL) AS directors
    FROM movies m
    LEFT JOIN movie_genre mg ON m.id = mg.movie_id
    LEFT JOIN genres g ON mg.genre_id = g.id
    LEFT JOIN movie_actors ma ON m.id = ma.movie_id
    LEFT JOIN actors a ON ma.actor_id = a.id
    LEFT JOIN movie_director md ON m.id = md.movie_id
    LEFT JOIN directors d ON md.director_id = d.id
	WHERE LOWER(m.title) LIKE LOWER($1)
	AND ($2 = '' OR LOWER(g.name) LIKE LOWER($2))
    GROUP BY m.id, m.title, m.synopsis, m.background, m.poster, m.release_date, m.duration, m.price
    ORDER BY m.id DESC
    LIMIT $3 OFFSET $4
    `

	searchPattern := "%" + search + "%"

	rows, err := conn.Query(context.Background(), baseQuery, searchPattern, genre, limit, offset)
	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows[Movies](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func GetMovieById(idMovie int) (Movies, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return Movies{}, err
	}
	defer conn.Conn().Close(context.Background())
	query := `
	SELECT 
		m.id,
		m.title,
		m.synopsis,
		m.background,
		m.poster,
		m.release_date,
		m.duration,
		m.price,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT g.name), NULL) AS genres,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT a.fullname), NULL) AS casts,
		ARRAY_REMOVE(ARRAY_AGG(DISTINCT d.fullname), NULL) AS directors
	FROM movies m
	LEFT JOIN movie_genre mg ON m.id = mg.movie_id
	LEFT JOIN genres g ON mg.genre_id = g.id
	LEFT JOIN movie_actors ma ON m.id = ma.movie_id
	LEFT JOIN actors a ON ma.actor_id = a.id
	LEFT JOIN movie_director md ON m.id = md.movie_id
	LEFT JOIN directors d ON md.director_id = d.id
	WHERE m.id = $1
	GROUP BY m.id, m.title, m.synopsis, m.background, m.poster, m.release_date, m.duration, m.price
	`
	rows, err := conn.Query(context.Background(), query, idMovie)
	if err != nil {
		return Movies{}, err
	}
	results, err := pgx.CollectOneRow[Movies](rows, pgx.RowToStructByName)
	return results, err

}

func InsertMovies(movie MoviesReq) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	if movie.Title == "" {
		return fmt.Errorf("Judul Tidak Boleh Kosong")
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var movieId int
	err = tx.QueryRow(context.Background(), `
		INSERT INTO movies (title, synopsis, background, poster, release_date, duration, price)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`, movie.Title, movie.Synopsis, movie.Background, movie.Poster, movie.ReleaseDate, movie.Duration, movie.Price).Scan(&movieId)

	if err != nil {
		return fmt.Errorf("Gagal insert movie: %v", err)
	}

	for _, genreId := range movie.Genres {
		_, err := tx.Exec(context.Background(), `
			INSERT INTO movie_genre (movie_id, genre_id) VALUES ($1, $2)
		`, movieId, genreId)
		if err != nil {
			return fmt.Errorf("Gagal insert movie_genre: %v", err)
		}
	}

	for _, directorId := range movie.Directors {
		_, err := tx.Exec(context.Background(), `
			INSERT INTO movie_director (movie_id, director_id) VALUES ($1, $2)
		`, movieId, directorId)
		if err != nil {
			return fmt.Errorf("Gagal insert movie_director: %v", err)
		}
	}

	for _, actorId := range movie.Casts {
		_, err := tx.Exec(context.Background(), `
			INSERT INTO movie_actors (movie_id, actor_id) VALUES ($1, $2)
		`, movieId, actorId)
		if err != nil {
			return fmt.Errorf("Gagal insert movie_actors: %v", err)
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("Gagal commit transaction: %v", err)
	}

	return nil
}

func UpdateMovies(movie Movies, movieId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var oldMovie Movies
	queryGet := `
		SELECT title, synopsis, background, poster, release_date, duration, price
		FROM movies
		WHERE id = $1`
	err = conn.QueryRow(context.Background(), queryGet, movieId).Scan(
		&oldMovie.Title,
		&oldMovie.Synopsis,
		&oldMovie.Background,
		&oldMovie.Poster,
		&oldMovie.ReleaseDate,
		&oldMovie.Duration,
		&oldMovie.Price,
	)
	if err != nil {
		return fmt.Errorf("failed to get existing movie: %w", err)
	}

	if movie.Title == "" {
		movie.Title = oldMovie.Title
	}
	if movie.Synopsis == "" {
		movie.Synopsis = oldMovie.Synopsis
	}
	if movie.Background == "" {
		movie.Background = oldMovie.Background
	}
	if movie.Poster == "" {
		movie.Poster = oldMovie.Poster
	}
	if movie.ReleaseDate.IsZero() {
		movie.ReleaseDate = oldMovie.ReleaseDate
	}
	if movie.Duration == 0 {
		movie.Duration = oldMovie.Duration
	}
	if movie.Price == 0 {
		movie.Price = oldMovie.Price
	}

	query := `
		UPDATE movies 
		SET title = $1, synopsis = $2, background = $3, poster = $4, release_date = $5, duration = $6, price = $7
		WHERE id = $8`
	result, err := tx.Exec(context.Background(), query,
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
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("movie with id %d not found", movieId)
	}

	if _, err = tx.Exec(context.Background(), `DELETE FROM movie_genre WHERE movie_id = $1`, movieId); err != nil {
		return err
	}
	if _, err = tx.Exec(context.Background(), `DELETE FROM movie_actors WHERE movie_id = $1`, movieId); err != nil {
		return err
	}
	if _, err = tx.Exec(context.Background(), `DELETE FROM movie_director WHERE movie_id = $1`, movieId); err != nil {
		return err
	}

	for _, genreId := range movie.Genres {
		if _, err = tx.Exec(context.Background(), `INSERT INTO movie_genre (movie_id, genre_id) VALUES ($1, $2)`, movieId, genreId); err != nil {
			return err
		}
	}

	for _, actorId := range movie.Casts {
		if _, err = tx.Exec(context.Background(), `INSERT INTO movie_actors (movie_id, actor_id) VALUES ($1, $2)`, movieId, actorId); err != nil {
			return err
		}
	}

	for _, directorId := range movie.Directors {
		if _, err = tx.Exec(context.Background(), `INSERT INTO movie_director (movie_id, director_id) VALUES ($1, $2)`, movieId, directorId); err != nil {
			return err
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return err
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
