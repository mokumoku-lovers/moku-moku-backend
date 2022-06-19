package users

import (
	"crypto/sha256"
	"encoding/hex"
	"moku-moku/utils/errors"
	"regexp"
	"strings"
	"unicode"
)

type User struct {
	Id          int64     `json:"id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Biography   string    `json:"biography"`
	Birthday    string    `json:"birthday"`
	ProfilePic  string    `json:"profile_picture"`
	Points      int64     `json:"points"`
	DateCreated string    `json:"date_created"`
	Passwords   Passwords `json:"passwords"`
}

type Passwords struct {
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
	PasswordR   string `json:"password_r"`
}

// IsUserPassword validates the password from the request body with
// the stored user password
func (p *Passwords) IsUserPassword(userID int64) *errors.RestErr {
	// Get userID's password
	var password string
	var err *errors.RestErr
	if password, err = p.GetUserPassword(userID); err != nil {
		return err
	}

	// Encrypts the password with SHA256
	hashedPassword := sha256.Sum256([]byte(p.OldPassword))
	providedPassword := hex.EncodeToString(hashedPassword[:])

	// Check if the retrieved password is the same as the one provided in the request
	if equal := password != providedPassword; equal {
		return errors.BadRequest("invalid passwords")
	}

	return nil
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

	if user.Passwords.Password == "" || user.Passwords.PasswordR == "" {
		return errors.BadRequest("invalid password")
	}

	// Go regexp doesn't support Lookaround backtrack
	number, upper, special, space := false, false, false, false
	for _, c := range user.Passwords.Password {
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
	if !number || !upper || !special || len(user.Passwords.Password) < 8 || space {
		return errors.BadRequest("invalid password")
	}

	if user.Passwords.Password != user.Passwords.PasswordR {
		return errors.BadRequest("passwords do not match")
	}

	return nil
}

func (user *User) UsernameValidation() *errors.RestErr {
	if len(user.Username) > 50 {
		return errors.BadRequest("username cannot be longer than 50 characters")
	}
	return nil
}
