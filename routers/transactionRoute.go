package routers

import (
	"weeklytickits/controller"
	"weeklytickits/middleware"

	"github.com/gin-gonic/gin"
)

func transactionRoutes(r *gin.RouterGroup) {

	r.POST("/", middleware.AuthMiddleware(), controller.CreateTransaction)
	r.GET("/", middleware.AuthMiddleware(), controller.GetAllTransactions)

}
