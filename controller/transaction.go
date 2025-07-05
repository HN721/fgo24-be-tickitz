package controller

import (
	"net/http"
	"strconv"
	"time"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

type CreateTransactionRequest struct {
	PriceTotal      int                               `json:"priceTotal"`
	UserId          int                               `json:"userId"`
	MovieId         int                               `json:"movieId"`
	CinemaId        int                               `json:"cinemaId"`
	PaymentMethodId int                               `json:"paymentMethodId"`
	Details         []models.TransactionDetailRequest `json:"details"`
}

func CreateTransaction(ctx *gin.Context) {
	var req CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	transaction := models.Transaction{
		Time:            time.Now(),
		Date:            time.Now(),
		PriceTotal:      req.PriceTotal,
		UserId:          req.UserId,
		MovieId:         req.MovieId,
		CinemaId:        req.CinemaId,
		PaymentMethodId: req.PaymentMethodId,
	}

	err := models.CreateTransactionWithDetails(transaction, req.Details)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create transaction",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Transaction created successfully",
	})
}

func GetAllTransactions(ctx *gin.Context) {
	data, err := models.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get transactions",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success get all transactions",
		Results: data,
	})
}

func GetTransactionByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid ID parameter",
			Error:   err.Error(),
		})
		return
	}

	data, err := models.GetTransactionById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Transaction not found",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success get transaction",
		Results: data,
	})
}

func GetTransactionsByUserID(ctx *gin.Context) {
	userParam := ctx.Param("userId")
	userId, err := strconv.Atoi(userParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid User ID parameter",
			Error:   err.Error(),
		})
		return
	}

	data, err := models.GetTransactionsByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get transactions",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success get user transactions",
		Results: data,
	})
}
