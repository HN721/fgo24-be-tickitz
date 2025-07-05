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
	r.POST("/genre", middleware.AdminMiddleware(), controller.CreateGenres)
	r.PATCH("/genre/:id", middleware.AdminMiddleware(), controller.UpdateGenre)
	r.DELETE("/genre/:id", middleware.AdminMiddleware(), controller.DeleteGenre)
	// actor
	r.GET("/actor", controller.GetActors)
	r.POST("/actor", middleware.AdminMiddleware(), controller.CreateActor)
	r.PATCH("/actor/:id", middleware.AdminMiddleware(), controller.UpdateActor)
	r.DELETE("/actor/:id", middleware.AdminMiddleware(), controller.DeleteActor)
	// director
	r.GET("/director", controller.GetDirector)
	r.POST("/director", middleware.AdminMiddleware(), controller.CreateDirector)
	r.PATCH("/director/:id", middleware.AdminMiddleware(), controller.UpdateDirector)
	r.DELETE("/director/:id", middleware.AdminMiddleware(), controller.DeleteDirector)
}
