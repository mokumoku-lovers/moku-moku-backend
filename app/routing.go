package app

import (
	"moku-moku/controllers/users"
)

func addRoutes() {
	router.GET("/users/:user_id", users.GetUser)
}
