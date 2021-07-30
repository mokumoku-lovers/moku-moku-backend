package users

import (
	"fmt"
	"moku-moku/datasources/postgresql/users_db"
	"moku-moku/utils/errors"
)

//User Data Access Object
const (
	queryInsertUser = "INSERT INTO users(email, username, display_name, biography, birthday, password, profile_pic, points, date_created) VALUES (?,?,?,?,?,?,?,?,?)"
)

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

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	defer stmt.Close()

	return nil
}
