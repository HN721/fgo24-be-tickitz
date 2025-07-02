package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func movieRoute(r *gin.RouterGroup) {
	r.GET("", controller.GetMovies)
}
