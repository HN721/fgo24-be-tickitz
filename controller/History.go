package controller

import (
	"net/http"
	"strconv"
	"weeklytickits/dto"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

// GetHistory godoc
// @Summary Get all history transactions
// @Description Get list of all transaction histories
// @Tags History
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /history [get]
// @Security Token
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

// UpdateHistory godoc
// @Summary Update a history transaction status
// @Description Update status and note of a history transaction by ID
// @Tags History
// @Accept json
// @Produce json
// @Param id path int true "History ID"
// @Param history body dto.HistoryReq true "History Update Request"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /history/{id} [patch]
// @Security Token
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

// GetHistoryId godoc
// @Summary Get History by User ID
// @Description Retrieve transaction history by authenticated user ID
// @Tags History
// @Produce json
// @Success 200 {object} utils.Response{results=[]dto.HistoryReq}
// @Failure 400 {object} utils.Response
// @Security Token
// @Router /history/user  [get]
func GetHistoryUserId(ctx *gin.Context) {
	userId, _ := ctx.Get("userID")
	data, err := models.GetHistoryByIdUser(userId.(int))
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
		Message: "Ok",
		Results: data,
	})
}
