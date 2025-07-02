package main

import (
	"fmt"
	"os"
	"weeklytickits/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routers.CombineRouter(r)
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
