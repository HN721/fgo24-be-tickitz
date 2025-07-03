package main

import (
	"fmt"
	"os"
	"weeklytickits/routers"
	"weeklytickits/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	routers.CombineRouter(r)
	services.FetchAndSaveActor()
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
