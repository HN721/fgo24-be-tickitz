package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func historyRoute(r *gin.RouterGroup) {
	r.GET("", controller.GetHistory)
	r.PATCH("/:id", controller.UpdateHistory)

}
