package controller

import (
	"net/http"
	"strconv"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

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
