package app

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	router = gin.Default()
)

func StartApp() {
	corsMiddlewareConfig := cors.DefaultConfig()
	corsMiddlewareConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "access_token"}
	corsMiddlewareConfig.AllowAllOrigins = true
	router.Use(cors.New(corsMiddlewareConfig))

	addRoutes()

	err := router.RunTLS(":9000", "../../certs/cert.pem", "../../certs/key.pem")

	if err != nil {
		errMsg := fmt.Errorf("routing couldn't be set up")
		fmt.Printf(errMsg.Error())
		os.Exit(-1)
	}
}
