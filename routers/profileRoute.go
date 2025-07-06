package routers

import (
	"weeklytickits/controller"

	"github.com/gin-gonic/gin"
)

func profileRoute(r *gin.RouterGroup) {
	r.GET("", controller.GetUserProfile)
	r.PATCH("", controller.UpdateProfileByUserId)
}
