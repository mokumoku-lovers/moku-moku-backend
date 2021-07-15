package users

import (
	"moku-moku/utils/errors"
	"regexp"
	"strings"
)

type User struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Biography  string `json:"biography"`
	Birthday   string `json:"birthday"`
	Password   string `json:"-"`
	ProfilePic string `json:"profile_picture"`
	Points     int32  `json:"points"`
}

func (user *User) EmailValidation() *errors.RestErr {
	// Email sanitation
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.BadRequest("invalid email address")
	}

	// Email validation
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

	validMail := emailRegex.MatchString(user.Email)

	if !validMail {
		return errors.BadRequest("invalid email address")
	}

	return nil
}
