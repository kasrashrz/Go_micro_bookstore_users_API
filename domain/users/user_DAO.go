package users

import (
	"fmt"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
)

var(
	usersDB = make(map[int64]*User)
)

func (user User) Get() *errors.RestErr{
	result := usersDB[user.Id]
	if result == nil{
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	user.Firstname = result.Firstname
	user.Lastname = result.Lastname
	return nil
}

func (user User) Create() *errors.RestErr {
	return nil
}