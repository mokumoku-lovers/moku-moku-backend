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
	PasswordR  string `json:"-"`
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

func (user *User) PasswordValidation() *errors.RestErr {

	if user.Password == "" || user.PasswordR == "" {
		return errors.BadRequest("invalid password")
	}

	/*
	*	At least one upper case English Letter, (?=.*?[A-Z])
	*	At least one lower case English letter, (?=.*?[a-z])
	*	At least one digit, (?=.*?[0-9])
	*	No spaces allowed, (?!.* )
	*	At least one special character, (?=.*?[#?!@$%^&*-])
	*	Minimum eight in length .{8,} (with the anchors)
	 */
	PasswordRegex := regexp.MustCompile(`^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?!.* )(?=.*?[#?!@$%^&*-]).{8,}$`)

	if !PasswordRegex.MatchString(user.Password) || !PasswordRegex.MatchString(user.PasswordR) {
		return errors.BadRequest("invalid password")
	}

	if user.Password != user.PasswordR {
		return errors.BadRequest("passwords do not match")
	}

	return nil
}
