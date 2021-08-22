package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	router = gin.Default()
)

func StartApp() {
	addRoutes()

	err := router.Run(":9000")

	if err != nil {
		errMsg := fmt.Errorf("routing couldn't be set up")
		fmt.Printf(errMsg.Error())
		os.Exit(-1)
	}
}
