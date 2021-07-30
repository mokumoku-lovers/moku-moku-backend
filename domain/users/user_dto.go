package users

import (
	"moku-moku/utils/errors"
	"regexp"
	"strings"
	"unicode"
)

type User struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Biography   string `json:"biography"`
	Birthday    string `json:"birthday"`
	Password    string `json:"-"`
	PasswordR   string `json:"-"`
	ProfilePic  string `json:"profile_picture"`
	Points      int32  `json:"points"`
	DateCreated string `json:"date_created"`
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

	// Go regexp doesn't support Lookaround backtrack
	number, upper, special, space := false, false, false, false
	for _, c := range user.Password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c):
		case c == ' ':
			space = true
		}
	}

	/*
	*	At least one upper case English Letter
	*	At least one lower case English letter
	*	At least one digit
	*	No spaces allowed
	*	At least one special character
	*	Minimum eight in length
	 */
	if !number || !upper || !special || len(user.Password) < 8 || space {
		return errors.BadRequest("invalid password")
	}

	if user.Password != user.PasswordR {
		return errors.BadRequest("passwords do not match")
	}

	return nil
}
