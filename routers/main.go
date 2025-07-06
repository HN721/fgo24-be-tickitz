package routers

import (
	"net/http"
	"weeklytickits/docs"
	"weeklytickits/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineRouter(r *gin.Engine) {
	authRouter(r.Group("/auth"))
	profileRoute(r.Group("/profile", middleware.AuthMiddleware()))
	movieRoute(r.Group("/movie"))
	cinemaRouter(r.Group("/cinema", middleware.AdminMiddleware()))
	transactionRoutes(r.Group("/trx", middleware.AuthMiddleware()))
	paymentRouter(r.Group("/payment", middleware.AdminMiddleware()))
	historyRoute(r.Group("/history"))

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/docs/index.html")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
