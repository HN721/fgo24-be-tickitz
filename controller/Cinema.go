package controller

import (
	"net/http"
	"strconv"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

func GetAllCinemas(ctx *gin.Context) {
	data, err := models.GetAllCinemas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get cinemas",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success get all cinemas",
		Results: data,
	})
}
func GetCinemaByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid ID",
			Error:   err.Error(),
		})
		return
	}

	data, err := models.GetCinemaByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Cinema not found",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success get cinema",
		Results: data,
	})
}

func CreateCinema(ctx *gin.Context) {
	var cinema models.Cinema
	if err := ctx.ShouldBindJSON(&cinema); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	err := models.InsertCinema(cinema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create cinema",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Cinema created successfully",
	})
}

func UpdateCinema(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid ID",
			Error:   err.Error(),
		})
		return
	}

	var cinema models.Cinema
	if err := ctx.ShouldBindJSON(&cinema); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	err = models.UpdateCinema(id, cinema)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update cinema",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Cinema updated successfully",
	})
}

func DeleteCinema(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid ID",
			Error:   err.Error(),
		})
		return
	}

	err = models.DeleteCinema(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to delete cinema",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Cinema deleted successfully",
	})
}
