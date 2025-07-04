package controller

import (
	"net/http"
	"strconv"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

func GetPayments(ctx *gin.Context) {
	data, err := models.GetAllPayments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to get payment methods",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success get all payment methods",
		Results: data,
	})
}

func GetPaymentByID(ctx *gin.Context) {
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

	data, err := models.GetPaymentByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Payment method not found",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success get payment method",
		Results: data,
	})
}

func CreatePayment(ctx *gin.Context) {
	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	err := models.InsertPayment(payment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create payment method",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Payment method created successfully",
	})
}

func UpdatePayment(ctx *gin.Context) {
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

	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	err = models.UpdatePayment(id, payment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update payment method",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Payment method updated successfully",
	})
}

func DeletePayment(ctx *gin.Context) {
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

	err = models.DeletePayment(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to delete payment method",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Payment method deleted successfully",
	})
}
