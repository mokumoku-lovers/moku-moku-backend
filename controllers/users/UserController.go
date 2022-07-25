package users

import (
	"crypto/sha256"
	"encoding/hex"
	"moku-moku/domain/users"
	"moku-moku/services"
	"moku-moku/utils/errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mokumoku-lovers/moku-moku-oauth-go/oauth"
)

func GetUser(c *gin.Context) {
	authErr := oauth.AuthenticateRequest(c.Request)
	if authErr != nil {
		c.JSON(authErr.Status, authErr)
		return
	}
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
	authErr := oauth.AuthenticateRequest(c.Request)
	if authErr != nil {
		c.JSON(authErr.Status, authErr)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	var passwords users.Passwords

	// Parse JSON and map it to Password model
	if err := c.ShouldBindJSON(&passwords); err != nil {
		resErr := errors.BadRequest("invalid json body")
		c.JSON(http.StatusBadRequest, resErr)
		return
	}
	// Check if it's user's password, only that user has right to delete
	// her/his own user
	if err := passwords.IsUserPassword(userId); err != nil {
		resErr := errors.BadRequest("invalid json body")
		c.JSON(http.StatusBadRequest, resErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func UpdateUser(c *gin.Context) {
	authErr := oauth.AuthenticateRequest(c.Request)
	if authErr != nil {
		c.JSON(authErr.Status, authErr)
		return
	}

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
	authErr := oauth.AuthenticateRequest(c.Request)
	if authErr != nil {
		c.JSON(authErr.Status, authErr)
		return
	}

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
	authErr := oauth.AuthenticateRequest(c.Request)
	if authErr != nil {
		c.JSON(authErr.Status, authErr)
		return
	}
	//Parse userId
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	//Parse JSON and map it to Passwords model
	var passwords users.Passwords
	if err := c.ShouldBind(&passwords); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	var user users.User
	user.Id = userId
	user.Passwords = passwords

	//Update password
	_, err := services.UpdatePassword(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "password successfully changed")
}

func UploadUserProfilePic(c *gin.Context) {
	authErr := oauth.AuthenticateRequest(c.Request)
	if authErr != nil {
		c.JSON(authErr.Status, authErr)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	//get file
	file, fileErr := c.FormFile("file")
	if fileErr != nil {
		c.JSON(http.StatusBadRequest, errors.BadRequest("a file must be uploaded"))
		return
	}

	fileType := file.Header.Get("Content-Type")
	if fileType != "image/jpeg" {
		c.JSON(http.StatusBadRequest, errors.BadRequest("file must be of type image"))
		return
	}

	name := file.Filename
	hashedName := sha256.Sum256([]byte(name))
	hashedNameString := hex.EncodeToString(hashedName[:])

	//map to user model
	var user users.User
	user.Id = userId
	user.ProfilePic = hashedNameString

	//write file to basePath
	basePath := "/MokuMoku/profile_pics/"
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		//create directory
		os.MkdirAll(basePath, 0700)
	}
	saveErr := c.SaveUploadedFile(file, basePath+hashedNameString+".png")
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError("file could not be saved"))
	}

	//partial update user with profile pic hashedName
	_, err := services.UpdateUser(true, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "user profile pic uploaded")
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
