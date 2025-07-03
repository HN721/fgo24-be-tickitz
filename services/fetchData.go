package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DirectorResponse struct {
	Results []DirectorResult `json:"results"`
}

type DirectorResult struct {
	ID int `json:"id"`

	Name string `json:"name"`
}
type ActorResponse struct {
	Results []ActorResult `json:"results"`
}

type ActorResult struct {
	ID int `json:"id"`

	Name string `json:"name"`
}
type TMDbResponse struct {
	Results []MovieResult `json:"results"`
}
type GenreResponse struct {
	Genres []Genre `json:"genres"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type MovieResult struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	Backdrop    string `json:"backdrop_path"`
	Poster      string `json:"poster_path"`
	ReleaseDate string `json:"release_date"`
}

func GetMovieDuration(movieID int) (int, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d", movieID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MmVhOGUwYjY1MGIyMDJkMTRlYmI1MjI5ZGQwZmRmOSIsIm5iZiI6MTc0NzM3Njk3NC42OTUwMDAyLCJzdWIiOiI2ODI2ZGI0ZTkxMTY1ZjYzYmE2ZWZjODAiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.UVz4N682u9la2B2SkINIeIYfKNJm8lvWBUzLCrs-3Wo")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var detail struct {
		Runtime int `json:"runtime"`
	}
	err = json.Unmarshal(body, &detail)
	if err != nil {
		return 0, err
	}

	return detail.Runtime, nil
}
func getGenres(conn *pgxpool.Conn) ([]Genre, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name FROM genres")
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[Genre])
}

func getActors(conn *pgxpool.Conn) ([]ActorResult, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, fullname as name FROM actors")
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[ActorResult])
}

func getDirectors(conn *pgxpool.Conn) ([]DirectorResult, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, fullname as name FROM directors")
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[DirectorResult])
}

func FetchMovie() error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Release()

	url := "https://api.themoviedb.org/3/movie/upcoming?language=en-US&page=1"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MmVhOGUwYjY1MGIyMDJkMTRlYmI1MjI5ZGQwZmRmOSIsIm5iZiI6MTc0NzM3Njk3NC42OTUwMDAyLCJzdWIiOiI2ODI2ZGI0ZTkxMTY1ZjYzYmE2ZWZjODAiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.UVz4N682u9la2B2SkINIeIYfKNJm8lvWBUzLCrs-3Wo")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var response TMDbResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	genres, err := getGenres(conn)
	if err != nil {
		return err
	}
	actors, err := getActors(conn)
	if err != nil {
		return err
	}
	directors, err := getDirectors(conn)
	if err != nil {
		return err
	}

	for _, movie := range response.Results {
		duration, err := GetMovieDuration(movie.ID)
		if err != nil || duration <= 0 {
			duration = 120
		}

		releaseDate, _ := time.Parse("2006-01-02", movie.ReleaseDate)

		var movieID int
		query := `
		INSERT INTO movies (title, synopsis, background, poster, release_date, duration, price)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
		err = conn.QueryRow(context.Background(), query,
			movie.Title,
			movie.Overview,
			"https://image.tmdb.org/t/p/original"+movie.Backdrop,
			"https://image.tmdb.org/t/p/original"+movie.Poster,
			releaseDate,
			duration,
			50000,
		).Scan(&movieID)

		if err != nil {
			fmt.Printf("Gagal insert movie %s: %v\n", movie.Title, err)
			continue
		}
		fmt.Printf("Berhasil insert movie: %s (ID: %d)\n", movie.Title, movieID)

		for i, g := range genres {
			if i >= 2 {
				break
			}
			conn.Exec(context.Background(), "INSERT INTO movie_genre (movie_id, genre_id) VALUES ($1, $2)", movieID, g.ID)
		}

		for i, a := range actors {
			if i >= 3 {
				break
			}
			conn.Exec(context.Background(), "INSERT INTO movie_actors (movie_id, actor_id) VALUES ($1, $2)", movieID, a.ID)
		}

		if len(directors) > 0 {
			conn.Exec(context.Background(), "INSERT INTO movie_director (movie_id, director_id) VALUES ($1, $2)", movieID, directors[0].ID)
		}
	}

	return nil
}

func FetchGenres() ([]Genre, error) {
	url := "https://api.themoviedb.org/3/genre/movie/list?language=en"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MmVhOGUwYjY1MGIyMDJkMTRlYmI1MjI5ZGQwZmRmOSIsIm5iZiI6MTc0NzM3Njk3NC42OTUwMDAyLCJzdWIiOiI2ODI2ZGI0ZTkxMTY1ZjYzYmE2ZWZjODAiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.UVz4N682u9la2B2SkINIeIYfKNJm8lvWBUzLCrs-3Wo")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var response GenreResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	fmt.Println(response.Genres)
	return response.Genres, nil
}
func FetchAndSaveGenres() error {
	genres, err := FetchGenres()
	if err != nil {
		return err
	}

	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	for _, genre := range genres {
		query := `INSERT INTO genres (name) VALUES ($1)`
		_, err := conn.Exec(context.Background(), query, genre.Name)
		if err != nil {
			fmt.Printf("Gagal insert genre %s: %v\n", genre.Name, err)
		} else {
			fmt.Printf("Berhasil insert genre: %s\n", genre.Name)
		}
	}

	return nil
}
func FetchActors() ([]ActorResult, error) {
	url := "https://api.themoviedb.org/3/person/popular?language=en-US&page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MmVhOGUwYjY1MGIyMDJkMTRlYmI1MjI5ZGQwZmRmOSIsIm5iZiI6MTc0NzM3Njk3NC42OTUwMDAyLCJzdWIiOiI2ODI2ZGI0ZTkxMTY1ZjYzYmE2ZWZjODAiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.UVz4N682u9la2B2SkINIeIYfKNJm8lvWBUzLCrs-3Wo")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var response ActorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Actors: %+v\n", response.Results)

	return response.Results, nil
}

func FetchAndSaveActor() error {
	actors, err := FetchActors()
	if err != nil {
		return err
	}

	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	for _, actor := range actors {
		query := `INSERT INTO actors (fullname) VALUES ($1)`
		_, err := conn.Exec(context.Background(), query, actor.Name)
		if err != nil {
			fmt.Printf("Gagal insert actor %s: %v\n", actor.Name, err)
		} else {
			fmt.Printf("Berhasil insert actor: %s\n", actor.Name)
		}
	}

	return nil
}
func FetchDirectors() ([]DirectorResult, error) {
	url := "https://api.themoviedb.org/3/person/popular?language=en-US&page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MmVhOGUwYjY1MGIyMDJkMTRlYmI1MjI5ZGQwZmRmOSIsIm5iZiI6MTc0NzM3Njk3NC42OTUwMDAyLCJzdWIiOiI2ODI2ZGI0ZTkxMTY1ZjYzYmE2ZWZjODAiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.UVz4N682u9la2B2SkINIeIYfKNJm8lvWBUzLCrs-3Wo")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var response DirectorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Directors (People): %+v\n", response.Results)

	return response.Results, nil
}

func FetchAndSaveDirector() error {
	directors, err := FetchDirectors()
	if err != nil {
		return err
	}

	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	for _, director := range directors {
		query := `INSERT INTO directors (fullname) VALUES ($1)`
		_, err := conn.Exec(context.Background(), query, director.Name)
		if err != nil {
			fmt.Printf("Gagal insert director %s: %v\n", director.Name, err)
		} else {
			fmt.Printf("Berhasil insert director: %s\n", director.Name)
		}
	}

	return nil
}
