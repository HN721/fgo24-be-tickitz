package routers

import (
	"weeklytickits/controller"
	"weeklytickits/middleware"

	"github.com/gin-gonic/gin"
)

func cinemaRouter(r *gin.RouterGroup) {
	r.GET("", controller.GetAllCinemas)
	r.GET("/:id", controller.GetCinemaByID)
	r.POST("", middleware.AdminMiddleware(), controller.CreateCinema)
	r.PATCH("/:id", middleware.AdminMiddleware(), controller.UpdateCinema)
	r.DELETE("/:id", middleware.AdminMiddleware(), controller.DeleteCinema)
}
