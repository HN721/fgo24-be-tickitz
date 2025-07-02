package controller

import (
	"net/http"
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
