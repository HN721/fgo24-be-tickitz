package main

import (
	"fmt"
	"os"
	"weeklytickits/routers"

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

	routers.CombineRouter(r)
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
