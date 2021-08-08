package services

import (
	"moku-moku/domain/users"
	"moku-moku/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	// Call middleware to sanitize and check if the fields are correct
	if err := user.EmailValidation(); err != nil {
		return nil, err
	}

	if err := user.PasswordValidation(); err != nil {
		return nil, err
	}

	// DTO save user to DB
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func UpdateUser(partialUpdate bool, user users.User) (*users.User, *errors.RestErr) {
	//Get user from db
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	//if partialUpdate, verify all fields to find what must be updated
	if partialUpdate {
		if user.Email != "" {
			if err := user.EmailValidation(); err != nil {
				return nil, err
			} else {
				current.Email = user.Email
			}
		}
		if user.Username != "" {
			current.Username = user.Username
		}
		if user.DisplayName != "" {
			current.DisplayName = user.DisplayName
		}
		}
	} else { //fullUpdate, update all to info in current user
	// Call middleware to sanitize and check if the fields are correct
	if err := user.EmailValidation(); err != nil {
		return nil, err
	}
	if err := user.PasswordValidation(); err != nil {
		return nil, err
	}

	// DTO save user to DB
	if err := user.Update(); err != nil {
		return nil, err
	}

	return &user, nil


}