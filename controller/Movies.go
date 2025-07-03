package controller

import (
	"net/http"
	"strconv"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

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
