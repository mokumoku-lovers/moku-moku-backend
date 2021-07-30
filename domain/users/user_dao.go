package users

import (
	"context"
	"fmt"
	"moku-moku/datasources/postgresql/users_db"
	"moku-moku/utils/date_utils"
	"moku-moku/utils/errors"
	"moku-moku/utils/pg_utils"
)

//User Data Access Object
const (
	queryInsertUser = "INSERT INTO user_db.users(email, username, display_name, biography, birthday, password, profile_pic, points, date_created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
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
	var err error

	user.DateCreated = date_utils.GetNowString()
	// TODO: Encrypt password with SHA algorithm

	// TODO: Failed queries increments users ID!!
	_, err = users_db.Client.Exec(context.Background(), queryInsertUser,
		user.Email, user.Username, user.DisplayName, user.Biography, nil, user.Password, user.ProfilePic, 0, user.DateCreated)

	if err != nil {
		return pg_utils.ParseError(err, "error when trying to save user")
	}

	//user.Id = userId

	return nil
}
