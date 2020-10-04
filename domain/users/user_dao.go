package users

import (
	"fmt"

	"gihub.com/eugenenosenko/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Save() (*User, *errors.RestErr) {
	if usersDB[u.ID] != nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("user with ID %d already exists", u.ID))
	}
	usersDB[u.ID] = u

	return u, nil
}

func (u *User) Get() (*User, *errors.RestErr) {
	res, ok := usersDB[u.ID]

	if !ok {
		return nil, errors.NewBadRequestError(fmt.Sprintf("user %d not found", u.ID))
	}

	user := &User{
		res.ID,
		res.FirstName,
		res.LastName,
		res.Email,
		res.DateCreated,
	}

	return user, nil
}
