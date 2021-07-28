package users

import (
	"fmt"
	"moku-moku/utils/errors"
)

//User Data Access Object

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.Email = result.Email
	user.Username = result.Username
	user.DisplayName = result.DisplayName
	user.Biography = result.Biography
	user.Birthday = result.Birthday
	user.Password = result.Password
	user.PasswordR = result.PasswordR
	user.ProfilePic = result.ProfilePic
	user.Points = result.Points
	user.DateCreated = result.DateCreated
	return nil
}
