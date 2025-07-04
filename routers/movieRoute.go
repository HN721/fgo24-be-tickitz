package routers

import (
	"weeklytickits/controller"
	"weeklytickits/middleware"

	"github.com/gin-gonic/gin"
)

func movieRoute(r *gin.RouterGroup) {
	r.GET("/upcoming", controller.GetUpcomingMovies)
	r.GET("/now-showing", controller.GetNowShoinfMovies)

	r.GET("", controller.GetMovies)
	r.POST("", middleware.AuthMiddleware(), controller.CreateMovies)
	r.PATCH("/:id", middleware.AuthMiddleware(), controller.UpdateMovies)
	r.DELETE("/:id", middleware.AuthMiddleware(), controller.DeleteMovies)
	// genre
	r.GET("/genre", controller.GetGenre)
	r.POST("/genre", controller.CreateGenres)
	r.PATCH("/genre/:id", controller.UpdateGenre)
	r.DELETE("/genre/:id", controller.DeleteGenre)
	// actor
	r.GET("/actor", controller.GetActors)
	r.POST("/actor", controller.CreateActor)
	r.PATCH("/actor/:id", controller.UpdateActor)
	r.DELETE("/actor/:id", controller.DeleteActor)
	// director
	r.GET("/director", controller.GetDirector)
	r.POST("/director", controller.CreateDirector)
	r.PATCH("/director/:id", controller.UpdateDirector)
	r.DELETE("/director/:id", controller.DeleteDirector)
}
