package controller

import (
	"net/http"
	"strconv"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

// GetAllPayment godoc
// @Summary Get  Payment
// @Description Retrieve all Payment
// @Tags Payment
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Payment}
// @Failure 500 {object} utils.Response
// @Router /payment [get]
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

// GetPaymentByID godoc
// @Summary Get payment method by ID
// @Description Retrieve a specific payment method using its ID
// @Tags Payment
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} utils.Response{results=models.Payment}
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /payment/{id} [get]
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

// CreatePayment godoc
// @Summary Create a new payment method
// @Description Add a new payment method to the database
// @Tags Payment
// @Accept json
// @Produce json
// @Param payment body models.Payment true "Payment Data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /payment [post]
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

// UpdatePayment godoc
// @Summary Update an existing payment method
// @Description Update payment method details by ID
// @Tags Payment
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Param payment body models.Payment true "Updated Payment Data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /payment/{id} [put]
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

// DeletePayment godoc
// @Summary Delete a payment method
// @Description Delete a payment method from database by ID
// @Tags Payment
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /payment/{id} [delete]
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
