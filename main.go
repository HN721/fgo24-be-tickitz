package main

import (
	"fmt"
	"time"
	"weeklytickits/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           Movxtar API Documentation
// @version         1.0
// @description     This is a backend service for tickitz web app
// @SecurityDefinitions.ApiKey  Token
// @in header
// @name Authorization
// @Basepath /
func main() {
	godotenv.Load()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://146.190.102.54:9202"},
		AllowMethods:     []string{"POST", "PATCH", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Static("/uploads", "./uploads")

	routers.CombineRouter(r)
	r.Run(fmt.Sprintf(":8080"))
}
