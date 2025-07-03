package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"weeklytickits/utils"
)

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

func FetchMovie() error {
	url := "https://api.themoviedb.org/3/movie/upcoming?language=en-US&page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MmVhOGUwYjY1MGIyMDJkMTRlYmI1MjI5ZGQwZmRmOSIsIm5iZiI6MTc0NzM3Njk3NC42OTUwMDAyLCJzdWIiOiI2ODI2ZGI0ZTkxMTY1ZjYzYmE2ZWZjODAiLCJzY29wZXMiOlsiYXBpX3JlYWQiXSwidmVyc2lvbiI6MX0.UVz4N682u9la2B2SkINIeIYfKNJm8lvWBUzLCrs-3Wo")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var response TMDbResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	for _, movie := range response.Results {
		duration, err := GetMovieDuration(movie.ID)
		if err != nil || duration <= 0 {
			duration = 120
		}
		releaseDate, err := time.Parse("2006-01-02", movie.ReleaseDate)
		if err != nil {
			releaseDate = time.Now()
		}

		query := `
		INSERT INTO movies (title, synopsis, background, poster, release_date, duration, price)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

		_, err = conn.Exec(context.Background(), query,
			movie.Title,
			movie.Overview,
			"https://image.tmdb.org/t/p/original"+movie.Backdrop,
			"https://image.tmdb.org/t/p/original"+movie.Poster,
			releaseDate,
			duration,
			50000)

		if err != nil {
			fmt.Printf("Gagal insert: %s\n", err)
		} else {
			fmt.Printf("Berhasil insert: %s\n", movie.Title)
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
