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

func UpdateUser(userId int64, user users.User) (*users.User, *errors.RestErr) {
	user.Id = userId

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