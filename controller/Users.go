package controller

import (
	"net/http"
	"os"
	"time"
	"weeklytickits/dto"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateToken(username string, role string, id int) (string, error) {
	godotenv.Load()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"role":     role,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
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
		return
	}
	users := &models.Users{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	err = models.Register(*users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Cant make User",
		})
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "OK",
		Results: users,
	})
}
func Login(ctx *gin.Context) {
	var req dto.LoginRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Input",
		})
		return
	}
	users := &models.Users{
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := models.Login(*users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Something Error",
		})
		return
	}
	generateToken, err := CreateToken(result.Username, result.Role, result.UserID)
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "ok",
		Results: generateToken,
	})
}
