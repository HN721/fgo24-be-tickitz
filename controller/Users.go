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
		return "", err
	}
	return tokenString, nil
}

// Get User
// @Summary Get User
// @Description Retrieve all users
// @Tags Auth
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Users}
// @Failure 400 {object} utils.Response
// @Security Token
// @Router /auth [get]
func GetUser(ctx *gin.Context) {
	result, err := models.FindAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Something Wrong On Database",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: result,
	})
}

// Post User
// @Summary Post User
// @Description Retrieve Post users
// @Tags Auth
// @Produce json
// @Param Register body dto.RegisterResquest true "User Data"
// @Success 200 {object} utils.Response{results=[]models.Users}
// @Failure 400 {object} utils.Response
// @Router /auth/register [post]
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
	if input.Password != input.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Confirm",
		})
		return
	}
	users := models.Users{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	err = models.Register(users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Cant make User",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "OK",
		Results: users,
	})
}

// Login godoc
// @Summary User Login
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param Login body dto.LoginRequest true "Login Data"
// @Success 200 {object} utils.Response{results=models.Users, token=string}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /auth/login [post]
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
	users := models.Users{
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := models.Login(users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Something Error",
			Error:   err.Error(),
		})
		return
	}
	generateToken, err := CreateToken(result.Username, result.Role, result.UserID)
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "ok",
		Results: result,
		Token:   generateToken,
	})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change password using OTP and old password verification
// @Tags Auth
// @Accept json
// @Produce json
// @Param ChangePassword body dto.ChangePassword true "Change Password Data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /auth/reset [post]
func ChangePassword(ctx *gin.Context) {
	var req dto.ChangePassword
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Data",
		})
		return
	}
	err = models.ChangePassword(req.Email, req.OTP, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed change Password",
		})
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "OK",
	})
}

// ForgotPassword godoc
// @Summary Forgot password request
// @Description Send OTP to user's email for password reset
// @Tags Auth
// @Accept json
// @Produce json
// @Security Token
// @Param ForgotPassword body dto.ForgotPasswordRequest true "Forgot Password Data"
// @Success 200 {object} utils.Response{}
// @Failure 400 {object} utils.Response{}
// @Router /auth/forgot [post]
func ForgotPassword(ctx *gin.Context) {
	var req dto.ForgotPasswordRequest
	ctx.ShouldBind(&req)
	err := models.ForgetPassword(req.Email)

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
	})

}
