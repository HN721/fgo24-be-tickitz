package controller

import (
	"net/http"
	"strconv"
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
	data, err := models.GetUpcomingMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Gagal mengambil data film",
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

// GetAllMovies godoc
// @Summary Get NowShowing movies
// @Description Retrieve all movies
// @Tags Movies
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Movies}
// @Failure 500 {object} utils.Response
// @Router /movie/now-showing [get]
func GetNowShoinfMovies(ctx *gin.Context) {
	data, err := models.NowShowingMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Gagal mengambil data film",
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

// @Summary Get Upcoming movies
// @Description Retrieve all movies
// @Tags Movies
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Movies}
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie [get]
func GetMovies(ctx *gin.Context) {
	data, err := models.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Gagal mengambil data film",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Berhasil mengambil data film",
		Results: data,
	})

}

// @Summary Create
// @Description Admin create movies
// @Tags Admin
// @Produce json
// @Accept json
// @Param movie body models.Movies true "Movie Data"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /movie [post]
func CreateMovies(ctx *gin.Context) {
	var req models.Movies

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	err = models.InsertMovies(req)
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
	data, err := models.GenreMovies()
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
