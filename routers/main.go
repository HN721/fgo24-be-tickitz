package routers

import (
	"net/http"
	"weeklytickits/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineRouter(r *gin.Engine) {
	authRouter(r.Group("/auth"))
	movieRoute(r.Group("/movie"))
	cinemaRouter(r.Group("/cinema"))

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/docs/index.html")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
