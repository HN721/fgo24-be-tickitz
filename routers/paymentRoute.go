package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func paymentRouter(r *gin.RouterGroup) {
	r.GET("", controller.GetPayments)
	r.GET("/:id", controller.GetPaymentByID)
	r.POST("", controller.CreatePayment)
	r.PATCH("/:id", controller.UpdatePayment)
	r.DELETE("/:id", controller.DeleteCinema)
}
