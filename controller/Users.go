package controller

import (
	"net/http"
	"weeklytickits/dto"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	var input models.Users
	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Input",
		})
	}
	result, err := models.FindAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Something Wrong On Database",
		})
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: result,
	})
}

func Register(ctx *gin.Context) {
	var input dto.RegisterResquest

	err := ctx.ShouldBind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Input",
		})
	}
	users := &models.Users{
		Email:    input.Email,
		Password: input.Password,
	}
	err = models.Register(*users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Cant make User",
		})
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "OK",
		Results: users,
	})
}
