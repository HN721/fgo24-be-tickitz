package main

import (
	"fmt"
	"os"
	"weeklytickits/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	routers.CombineRouter(r)
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
