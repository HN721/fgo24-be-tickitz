package controller

import (
	"fmt"
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
	Location        string                            `json:"location"`
	MovieId         int                               `json:"movieId"`
	CinemaId        int                               `json:"cinemaId"`
	PaymentMethodId int                               `json:"paymentMethodId"`
	Details         []models.TransactionDetailRequest `json:"details"`
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a transaction along with its details
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body controller.CreateTransactionRequest true "Transaction Data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /trx [post]
func CreateTransaction(ctx *gin.Context) {
	var req CreateTransactionRequest
	userId, _ := ctx.Get("userID")
	fmt.Println(userId)
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
		Location:        req.Location,
		UserId:          userId.(int),
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

// GetAllTransactions godoc
// @Summary Get all transactions
// @Description Retrieve all transactions from database
// @Tags Transaction
// @Produce json
// @Success 200 {object} utils.Response{results=[]dto.TransactionResponses}
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /trx [get]
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

// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Description Retrieve a specific transaction using its ID
// @Tags Transaction
// @Produce json
// @Param id path int true "Transaction ID"
// @Security Token
// @Success 200 {object} utils.Response{results=dto.TransactionResponses}
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /trx/{id} [get]
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

// GetTransactionsByUserID godoc
// @Summary Get all transactions by user ID
// @Description Retrieve all transactions for a specific user
// @Tags Transaction
// @Produce json
// @Param id path int true "User ID"
// @Security Token
// @Success 200 {object} utils.Response{results=[]dto.TransactionResponses}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /trx/user/{id} [get]
func GetTransactionsByUserID(ctx *gin.Context) {
	userParam := ctx.Param("id")
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

// GetTransactionDetail godoc
// @Summary Get transaction details
// @Description Retrieve detailed information of a transaction by transaction ID
// @Tags Transaction
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} utils.Response{results=[]dto.TransactionDetailData}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /trx/detail/{id} [get]
func GetTransactionDetail(ctx *gin.Context) {
	transactionIDStr := ctx.Param("id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid transaction ID",
		})
		return
	}

	details, err := models.GetTransactionDetailWithInfoByTransactionId(transactionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get transaction detail",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: details,
	})
}
