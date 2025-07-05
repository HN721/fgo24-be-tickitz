package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func transactionRoutes(r *gin.RouterGroup) {

	r.POST("/", controller.CreateTransaction)
	r.GET("/", controller.GetAllTransactions)
	r.GET("/:id", controller.GetTransactionByID)
	r.GET("/user/:id", controller.GetTransactionsByUserID)
	r.GET("/detail/:id", controller.GetTransactionDetail)

}
