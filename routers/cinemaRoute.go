package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func cinemaRouter(r *gin.RouterGroup) {
	r.GET("", controller.GetAllCinemas)
	r.GET("/:id", controller.GetCinemaByID)
	r.POST("", controller.CreateCinema)
	r.PATCH("/:id", controller.UpdateCinema)
	r.DELETE("/:id", controller.DeleteCinema)
}
