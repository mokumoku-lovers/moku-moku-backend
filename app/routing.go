package app

import (
	"moku-moku/controllers/users"
)

func addRoutes() {
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.PUT("/users/:user_id", users.UpdateUser)
}
