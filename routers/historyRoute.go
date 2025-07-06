package routers

import (
	"weeklytickits/controller"
	"weeklytickits/middleware"

	"github.com/gin-gonic/gin"
)

func historyRoute(r *gin.RouterGroup) {
	r.GET("", middleware.AdminMiddleware(), controller.GetHistory)
	r.GET("user", middleware.AuthMiddleware(), controller.GetHistoryUserId)
	r.PATCH("/:id", middleware.AdminMiddleware(), controller.UpdateHistory)

}
