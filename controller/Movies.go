package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

// GetAllMovies godoc
// @Summary Get Upcoming movies
// @Description Retrieve all movies
// @Tags Movies
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Movies}
// @Failure 500 {object} utils.Response
// @Router /movie/upcoming [get]
func GetUpcomingMovies(ctx *gin.Context) {
	cacheKey := "upcoming_movies"

	cached, err := utils.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil && cached != "" {
		ctx.Data(http.StatusOK, "application/json", []byte(cached))
		return
	}

	data, err := models.GetUpcomingMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Gagal mengambil data film",
			Error:   err.Error(),
		})
		return
	}

	response := utils.Response{
		Success: true,
		Message: "OK",
		Results: data,
	}

	jsonData, _ := json.Marshal(response)
	utils.RedisClient.Set(context.Background(), cacheKey, jsonData, 5*time.Minute)

	ctx.JSON(http.StatusOK, response)
}

// @Summary Get Movies By Genre
// @Description Retrieve movies filtered by genre
// @Tags Movies
// @Produce json
// @Param genre query string true "Genre Name"
// @Success 200 {object} utils.Response{results=[]models.Movies}
// @Failure 500 {object} utils.Response
// @Router /movie/filter [get]
func GetFilterMovie(ctx *gin.Context) {
	genre := ctx.Query("genre")
	if genre == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Genre query parameter is required",
		})
		return
	}

	movies, err := models.FilterMoviesByGenre(genre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error retrieving movies by genre",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Successfully retrieved movies by genre",
		Results: movies,
	})
}

// GetAllMovies godoc
// @Summary Get NowShowing movies
// @Description Retrieve all movies
// @Tags Movies
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Movies}
// @Failure 500 {object} utils.Response
// @Router /movie/now-showing [get]
func GetNowShoinfMovies(ctx *gin.Context) {
	cacheKey := ctx.Request.RequestURI

	// Cek apakah data ada di Redis
	cachedData, err := utils.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil && cachedData != "" {
		// Jika ditemukan di cache, langsung kirim data dari Redis
		ctx.Data(http.StatusOK, "application/json", []byte(cachedData))
		return
	}

	// Jika tidak ada di cache, ambil dari database
	data, err := models.NowShowingMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Gagal mengambil data film",
			Error:   err.Error(),
		})
		return
	}

	// Buat JSON response untuk disimpan di Redis
	response := utils.Response{
		Success: true,
		Message: "OK",
		Results: data,
	}
	jsonData, err := json.Marshal(response)
	if err == nil {
		// Simpan ke Redis selama 5 menit
		utils.RedisClient.Set(context.Background(), cacheKey, jsonData, 5*time.Minute)
	}

	// Kembalikan data ke client
	ctx.JSON(http.StatusOK, response)
}

// @Summary Get One Movie
// @Description Get Detail Movie
// @Tags Movies
// @Produce json
// @Accept json
// @Param id path int true "Movie ID"
// @Success 200 {object} utils.Response{results=[]models.Movies}
// @Failure 500 {object} utils.Response
// @Router /movie/detail/{id} [get]
func GetMovieById(ctx *gin.Context) {
	movieId := ctx.Param("id")
	id, _ := strconv.Atoi(movieId)

	result, err := models.GetMovieById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Failed to get Movie",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Sucessfully Get Movie Data",
		Results: result,
	})
}

// @Summary Get Upcoming movies
// @Description Retrieve all movies with search and pagination
// @Tags Movies
// @Produce json
// @Param genre query string false "genre by title"
// @Param search query string false "Search by title"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(5)
// @Success 200 {object} utils.Response{results=[]models.Movies}
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie [get]
func GetMovies(ctx *gin.Context) {
	cacheKey := ctx.Request.RequestURI

	// Cek apakah ada data di Redis
	cached := utils.RedisClient.Get(context.Background(), cacheKey)
	if cached.Err() == nil {
		// Jika data ada di Redis, kembalikan langsung ke client
		ctx.Data(http.StatusOK, "application/json", []byte(cached.Val()))
		return
	}

	// Ambil query params
	search := ctx.Query("search")
	genre := ctx.Query("genre")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Ambil data dari database
	movies, err := models.GetAllMovies(search, genre, pageInt, limitInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error retrieving movies",
			Error:   err.Error(),
		})
		return
	}

	// Encode response ke JSON string untuk simpan di Redis
	response := utils.Response{
		Success: true,
		Message: "Success get all movies",
		Results: movies,
	}
	jsonBytes, err := json.Marshal(response)
	if err == nil {
		// Simpan ke Redis dengan TTL (misalnya 5 menit)
		utils.RedisClient.Set(context.Background(), cacheKey, string(jsonBytes), 5*time.Minute)
	}

	// Kembalikan response ke client
	ctx.JSON(http.StatusOK, response)
}

// @Summary Create
// @Description Admin create movies
// @Tags Admin
// @Produce json
// @Accept json
// @Param movie body models.MoviesReq true "Movie Data"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie [post]
func CreateMovies(ctx *gin.Context) {
	var req models.MoviesReq

	posterFile, posterErr := ctx.FormFile("poster")
	backgroundFile, backgroundErr := ctx.FormFile("background")

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		uploadPath = "uploads/movies"
	}

	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, os.ModePerm)
	}

	if posterErr == nil && posterFile != nil {
		posterName := fmt.Sprintf("poster-%d-%s", time.Now().Unix(), posterFile.Filename)
		savePath := filepath.Join(uploadPath, posterName)

		if err := ctx.SaveUploadedFile(posterFile, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.Response{
				Success: false,
				Message: "Failed to upload poster",
				Error:   err.Error(),
			})
			return
		}
		req.Poster = posterName
	}

	if backgroundErr == nil && backgroundFile != nil {
		bgName := fmt.Sprintf("bg-%d-%s", time.Now().Unix(), backgroundFile.Filename)
		savePath := filepath.Join(uploadPath, bgName)

		if err := ctx.SaveUploadedFile(backgroundFile, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.Response{
				Success: false,
				Message: "Failed to upload background",
				Error:   err.Error(),
			})
			return
		}
		req.Background = bgName
	}

	err := models.InsertMovies(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "OK",
		Results: req,
	})
}

// @Summary Update
// @Description Admin Update movies
// @Tags Admin
// @Produce json
// @Accept json
// @Param id path int true "Movie ID"
// @Param movie body models.Movies true "Movie Data"
// @Success 200 {object} utils.Response "Successfully updated movie"
// @Failure 400 {object} utils.Response "Bad Request"
// @Failure 500 {object} utils.Response "Internal Server Error"
// @Security Token
// @Router /movie/{id} [patch]
func UpdateMovies(ctx *gin.Context) {
	var req models.Movies
	id := ctx.Param("id")
	movieId, _ := strconv.Atoi(id)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error While INput",
			Error:   err.Error(),
		})
		return
	}

	err = models.UpdateMovies(req, movieId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error in Database",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, utils.Response{
		Success: true,
		Message: "OK",
		Results: req,
	})
}

// @Summary Delete
// @Description Delete Update movies
// @Tags Admin
// @Produce json
// @Accept json
// @Param id path int true "Movie ID"
// @Param movie body models.Movies true "Movie Data"
// @Success 200 {object} utils.Response "Successfully updated movie"
// @Failure 400 {object} utils.Response "Bad Request"
// @Failure 500 {object} utils.Response "Internal Server Error"
// @Security Token
// @Router /movie/{id} [delete]
func DeleteMovies(ctx *gin.Context) {
	id := ctx.Param("id")
	movieId, _ := strconv.Atoi(id)

	err := models.DeleteMovies(movieId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Cannot delete Movie",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, utils.Response{
		Success: true,
		Message: "OK",
	})
}

// Movies Genres
// @Summary Get Genre
// @Description Retrieve all Genre
// @Tags Genres
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Genres}
// @Failure 500 {object} utils.Response
// @Router /movie/genre [get]
func GetGenre(ctx *gin.Context) {
	cacheKey := "genres"

	// Coba ambil dari cache
	cachedData, err := utils.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil && cachedData != "" {
		ctx.Data(http.StatusOK, "application/json", []byte(cachedData))
		return
	}

	// Ambil dari database jika tidak ada di Redis
	data, err := models.GenreMovies()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Gagal mengambil genre",
			Error:   err.Error(),
		})
		return
	}

	// Simpan ke Redis
	response := utils.Response{
		Success: true,
		Message: "OK",
		Results: data,
	}
	jsonData, _ := json.Marshal(response)
	utils.RedisClient.Set(context.Background(), cacheKey, jsonData, 5*time.Minute)

	ctx.JSON(http.StatusOK, response)
}

// @Summary Create Genre
// @Description Admin create Genre
// @Tags Genres
// @Produce json
// @Accept json
// @Param movie body models.Genres true "Genre Data"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/genre [post]
func CreateGenres(ctx *gin.Context) {
	var req models.Genres

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	err = models.CreateGenre(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "OK",
		Results: req,
	})
}

// @Summary Update Genre
// @Description Admin Update Genre
// @Tags Genres
// @Produce json
// @Accept json
// @Param movie body models.Genres true "Genre Data"
// @Param id path int true "Genre ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/genre/{id} [Patch]
func UpdateGenre(ctx *gin.Context) {
	var req models.Genres
	id := ctx.Param("id")
	genreId, _ := strconv.Atoi(id)
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	err = models.UpdateGenre(req, genreId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: req,
	})
}

// @Summary Delete Genre
// @Description Admin Delete Genre
// @Tags Genres
// @Produce json
// @Accept json
// @Param movie body models.Genres true "Genre Data"
// @Param id path int true "Genre ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/genre/{id} [Delete]
func DeleteGenre(ctx *gin.Context) {
	id := ctx.Param("id")
	genreId, _ := strconv.Atoi(id)
	err := models.DeleteGenre(genreId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
	})
}

// Movies Actors
// @Summary Get Actorss
// @Description Retrieve all actors
// @Tags Actors
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Actor}
// @Failure 400 {object} utils.Response
// @Router /movie/actor [get]
func GetActors(ctx *gin.Context) {
	data, err := models.ActorMovies()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: data,
	})
}

// @Summary Create Actors
// @Description Admin create Actors
// @Tags Actors
// @Produce json
// @Accept json
// @Param movie body models.Actor true "Actor Data"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/actor [post]
func CreateActor(ctx *gin.Context) {
	var req models.Actor

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	err = models.CreateActor(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "OK",
		Results: req,
	})
}

// @Summary Update Actor
// @Description Admin update Actors
// @Tags Actors
// @Produce json
// @Accept json
// @Param movie body models.Actor true "Actor Data"
// @Param id path int true "Actor ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/actor/{id} [patch]
func UpdateActor(ctx *gin.Context) {
	var req models.Actor
	id := ctx.Param("id")
	actorId, _ := strconv.Atoi(id)
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	err = models.UpdateActor(req, actorId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: req,
	})
}

// @Summary Delete Actor
// @Description Admin Delete Actors
// @Tags Actors
// @Produce json
// @Accept json
// @Param id path int true "Actor ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/actor/{id} [delete]
func DeleteActor(ctx *gin.Context) {
	id := ctx.Param("id")
	actorId, _ := strconv.Atoi(id)
	err := models.DeleteActor(actorId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
	})
}

// Movies Director
// Movies Director
// @Summary Get Director
// @Description Retrieve all director
// @Tags Directors
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Directors}
// @Failure 400 {object} utils.Response
// @Router /movie/director [get]
func GetDirector(ctx *gin.Context) {
	data, err := models.DirectorsMovie()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: data,
	})
}

// @Summary Create Directors
// @Description Admin create Directors
// @Tags Directors
// @Produce json
// @Accept json
// @Param movie body models.Directors true "Director Data"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/director [post]
func CreateDirector(ctx *gin.Context) {
	var req models.Directors

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	err = models.CreateDirector(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "OK",
		Results: req,
	})
}

// @Summary Update Directors
// @Description Admin Update Directors
// @Tags Directors
// @Produce json
// @Accept json
// @Param movie body models.Directors true "Director Data"
// @Param id path int true "Director ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/director/{id} [patch]
func UpdateDirector(ctx *gin.Context) {
	var req models.Directors
	id := ctx.Param("id")
	directorId, _ := strconv.Atoi(id)
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	err = models.UpdateDirector(req, directorId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: req,
	})
}

// @Summary Delete Directors
// @Description Admin Delete Directors
// @Tags Directors
// @Produce json
// @Accept json
// @Param id path int true "Director ID"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie/director/{id} [delete]
func DeleteDirector(ctx *gin.Context) {
	id := ctx.Param("id")
	directorId, _ := strconv.Atoi(id)
	err := models.DeleteDirector(directorId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
	})
}
