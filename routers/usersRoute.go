package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/pass", controller.ChangePassword)
}
