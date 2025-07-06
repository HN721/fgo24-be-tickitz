package routers

import (
	"weeklytickits/controller"
	"weeklytickits/middleware"

	"github.com/gin-gonic/gin"
)

func paymentRouter(r *gin.RouterGroup) {
	r.GET("", controller.GetPayments)
	r.GET("/:id", controller.GetPaymentByID)
	r.POST("", middleware.AdminMiddleware(), controller.CreatePayment)
	r.PATCH("/:id", middleware.AdminMiddleware(), controller.UpdatePayment)
	r.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteCinema)
}
