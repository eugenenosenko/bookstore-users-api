package services

import (
	"fmt"

	"gihub.com/eugenenosenko/bookstore-users-api/domain/users"
	"gihub.com/eugenenosenko/bookstore-users-api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := users.User{ID: userID}

	result, err := user.Get()

	if err != nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("user with id %d does not exist", userID))
	}

	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	result, err := user.Save()

	if err != nil {
		return nil, err
	}

	return result, nil
}
