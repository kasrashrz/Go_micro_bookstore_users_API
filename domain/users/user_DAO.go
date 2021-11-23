package users

import (
	"fmt"
	. "github.com/kasrashrz/Golang_microservice/datastore/mysql/user_db"
	"github.com/kasrashrz/Golang_microservice/utils/dates"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
)

const (

)

var (
	UDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	result := UDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	user.Firstname = result.Firstname
	user.Lastname = result.Lastname

	return nil
}

func (user *User) Create() *errors.RestErr {
	currentUser := UDB[user.Id]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.BadRequest(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.BadRequest(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = dates.GetCurrentTimeString()

	UDB[user.Id] = user

	return nil
}
