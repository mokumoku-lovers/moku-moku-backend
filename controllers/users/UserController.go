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
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
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

	c.JSON(http.StatusOK, createdUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func UpdateUser(c *gin.Context) {
	// Parse userId
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	// Parse JSON and map it to User model
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user.Id = userId

	//Check if patch request
	partialUpdate := c.Request.Method == http.MethodPatch

	// Update user
	updatedUser, err := services.UpdateUser(partialUpdate, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, updatedUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUserPoints(c *gin.Context) {
	// Parse userId & userPoints
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	userPoints, pointErr := strconv.ParseInt(c.Param("user_points"), 10, 64)
	if pointErr != nil {
		err := errors.BadRequest("user_points should be a number")
		c.JSON(err.Status, err)
		return
	}

	// Parse JSON and map it to User model
	var user users.User
	user.Id = userId
	user.Points = userPoints

	// Update user
	updatedUser, err := services.UpdateUser(true, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, updatedUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUserPassword(c *gin.Context) {
	//Parse userId
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	//Parse JSON and map it to User model
	var user users.User
	if err := c.ShouldBind(&user); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	user.Id = userId

	//Update password
	_, err := services.UpdatePassword(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "password successfully changed")
}

func Login(c *gin.Context) {
	var request users.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	user, err := services.LoginUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
