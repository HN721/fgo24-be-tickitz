package controller

import (
	"net/http"
	"strconv"
	"weeklytickits/dto"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

func GetHistory(ctx *gin.Context) {
	data, err := models.GetHistory()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: data,
	})
}
func UpdateHistory(ctx *gin.Context) {
	historyId := ctx.Param("id")
	id, _ := strconv.Atoi(historyId)
	var req dto.HistoryReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Error",
			Error:   err.Error(),
		})
		return
	}
	err = models.UpdateHistory(id, req)
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
		Message: "Sucess Update",
	})
}
