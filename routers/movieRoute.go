package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func movieRoute(r *gin.RouterGroup) {
	r.GET("", controller.GetMovies)
	r.POST("", controller.CreateMovies)
	r.PATCH("/:id", controller.UpdateMovies)
	r.DELETE("/:id", controller.DeleteMovies)
	// genre
	r.GET("/genre", controller.GetGenre)
	r.POST("/genre", controller.CreateGenres)
	r.PATCH("/genre", controller.UpdateGenre)

}
