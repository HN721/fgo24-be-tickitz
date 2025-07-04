package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	r.GET("", controller.GetUser)
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/reset", controller.ChangePassword)
	r.POST("/forgot", controller.ForgotPassword)
}
