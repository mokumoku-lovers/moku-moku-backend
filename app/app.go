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
	addRoutes()

	router.Use(cors.Default())
	err := router.Run(":9000")

	if err != nil {
		errMsg := fmt.Errorf("routing couldn't be set up")
		fmt.Printf(errMsg.Error())
		os.Exit(-1)
	}
}
