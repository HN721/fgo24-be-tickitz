package controller

import (
	"net/http"
	"strconv"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

// GetAllCinema godoc
// @Summary Get  Cinema
// @Description Retrieve all Cinema
// @Tags Cinema
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Cinema}
// @Failure 500 {object} utils.Response
// @Router /cinema [get]
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

// GetCinemaByID godoc
// @Summary Get cinema by ID
// @Description Retrieve a specific cinema using its ID
// @Tags Cinema
// @Produce json
// @Param id path int true "Cinema ID"
// @Success 200 {object} utils.Response{results=models.Cinema}
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /cinema/{id} [get]
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

// CreateCinema godoc
// @Summary Create a new cinema
// @Description Add a new cinema to the database
// @Tags Cinema
// @Accept json
// @Produce json
// @Security Token
// @Param cinema body models.Cinema true "Cinema Data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /cinema [post]
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

// UpdateCinema godoc
// @Summary Update an existing cinema
// @Description Update cinema details by ID
// @Tags Cinema
// @Accept json
// @Produce json
// @Security Token
// @Param id path int true "Cinema ID"
// @Param cinema body models.Cinema true "Updated Cinema Data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /cinema/{id} [put]
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

// DeleteCinema godoc
// @Summary Delete a cinema
// @Description Delete a cinema from database by ID
// @Tags Cinema
// @Produce json
// @Param id path int true "Cinema ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /cinema/{id} [delete]
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
