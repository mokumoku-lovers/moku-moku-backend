package users

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"moku-moku/datasources/postgresql/users_db"
	"moku-moku/utils/date_utils"
	"moku-moku/utils/errors"
	"moku-moku/utils/pg_utils"
	"strings"
	"time"

	"github.com/georgysavva/scany/pgxscan"
)

//User Data Access Object
const (
	queryInsertUser = "INSERT INTO user_db.users(email, username, display_name, biography, birthday, password, profile_pic, points, date_created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;"
	queryGetUser    = "SELECT id, email, username, display_name, biography, COALESCE(to_char(birthday, 'YYYY-MM-DD'), '') AS birthday, password, profile_pic, points, to_char(date_created, 'YYYY-MM-DD') AS date_created FROM user_db.users WHERE id =$1;"
	queryDeleteUser = "DELETE FROM user_db.users WHERE id = $1;"
	queryUpdateUser = "UPDATE user_db.users SET email=$2, username=$3, display_name=$4, biography=$5, birthday=$6, password=$7, profile_pic=$8, points=$9 WHERE id=$1;"
)

func (user *User) Get() *errors.RestErr {
	//err := users_db.Client.QueryRow(context.Background(), queryGetUser, user.Id).Scan(&user.Id, &user.Email, &user.Username, &user.DisplayName, &user.Biography, &user.Birthday, &user.Password, &user.ProfilePic, &user.Points, &user.DateCreated)
	var users []*User
	err := pgxscan.Select(context.Background(), users_db.Client, &users, queryGetUser, user.Id)
	if err != nil {
		return pg_utils.ParseError(err, "error when trying to get user")
	}
	*user = *users[0]
	return nil
}

func (user *User) Save() *errors.RestErr {
	var err error

	// Dates
	user.DateCreated = date_utils.GetNowString()
	var birthday time.Time
	if user.Birthday != "" {
		birthday, err = time.Parse(date_utils.DateFormat, user.Birthday)
	}

	if err == nil {
		user.Birthday = strings.Fields(birthday.String())[0]
	}

	// Encrypts the password with SHA256
	hashedPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashedPassword[:])

	// TODO: Failed queries increments users ID!!
	var stmt pgx.Row
	if user.Birthday != "" {
		stmt = users_db.Client.QueryRow(context.Background(), queryInsertUser,
			user.Email, user.Username, user.DisplayName, user.Biography, user.Birthday, user.Password, user.ProfilePic, 0, user.DateCreated)
	} else {
		stmt = users_db.Client.QueryRow(context.Background(), queryInsertUser,
			user.Email, user.Username, user.DisplayName, user.Biography, nil, user.Password, user.ProfilePic, 0, user.DateCreated)
	}

	err = stmt.Scan(&user.Id)
	if err != nil {
		return pg_utils.ParseError(err, "error when trying to save user")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Exec(context.Background(), queryDeleteUser, user.Id)
	if err != nil {
		return pg_utils.ParseError(err, "error when trying to delete user")
	}
	if stmt.RowsAffected() != 1 {
		return errors.NotFoundError("user does not exist")
	}

	return nil
}

func (user *User) Update() *errors.RestErr {
	// Parse Birthday
	if user.Birthday != "" {
		birthday, _ := time.Parse(date_utils.DateFormat, user.Birthday)
		user.Birthday = strings.Fields(birthday.String())[0]
	}

	// Encrypts the password with SHA256
	// TODO: If password is not changed do not re-hash the hash
	hashedPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashedPassword[:])

	var stmt pgconn.CommandTag
	var err error
	if user.Birthday != "" {
		stmt, err = users_db.Client.Exec(context.Background(), queryUpdateUser, user.Id, user.Email, user.Username, user.DisplayName, user.Biography, user.Birthday, user.Password, user.ProfilePic, user.Points)
	} else {
		stmt, err = users_db.Client.Exec(context.Background(), queryUpdateUser, user.Id, user.Email, user.Username, user.DisplayName, user.Biography, nil, user.Password, user.ProfilePic, user.Points)
	}

	if err != nil {
		return pg_utils.ParseError(err, "error when trying to update user")
	}
	if stmt.RowsAffected() != 1 {
		return errors.NotFoundError("user does not exist")
	}

	return nil
}
