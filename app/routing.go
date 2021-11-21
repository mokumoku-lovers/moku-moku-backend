package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"moku-moku/controllers/users"
	"net/http"
)

func addRoutes() {
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id/points/:user_points", users.UpdateUserPoints)
	router.PATCH("/users/:user_id/change_password", users.UpdateUserPassword)
	router.POST("/users/login", users.Login)

	// Swagger Documentation
	opts := middleware.RedocOpts{SpecURL: "./swagger.yml", Title: "Moku-Moku-Users"}
	swg := middleware.Redoc(opts, nil)

	router.GET("/docs", gin.WrapH(swg))
	router.GET("/swagger.yml", gin.WrapH(http.FileServer(http.Dir("./"))))
}
