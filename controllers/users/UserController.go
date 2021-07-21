package users

import (
	"moku-moku/domain/users"
	"moku-moku/services"
	"moku-moku/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	// Parse JSON and map it to User model
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	// Create a user
	createdUser, err := services.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
}
