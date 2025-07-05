package routers

import (
	"weeklytickits/controller"
	"weeklytickits/middleware"

	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	r.GET("", middleware.AdminMiddleware(), controller.GetUser)
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/reset", middleware.AuthMiddleware(), controller.ChangePassword)
	r.POST("/forgot", middleware.AuthMiddleware(), controller.ForgotPassword)
}
