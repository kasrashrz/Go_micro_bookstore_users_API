package users

import (
	"fmt"
	. "github.com/kasrashrz/Golang_microservice/datastore/mysql/user_db"
	"github.com/kasrashrz/Golang_microservice/utils/dates"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
)

const (
	//queryInsertUser = "INSERT INTO users_db.users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

func (user *User) Get() *errors.RestErr {
	if err := Client.Ping(); err != nil {
		panic(err)
	}

	//result := UDB[user.Id]
	//if result == nil {
	//	return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	//}
	//user.Id = result.Id
	//user.Email = result.Email
	//user.DateCreated = result.DateCreated
	//user.Firstname = result.Firstname
	//user.Lastname = result.Lastname

	return nil
}

func (user *User) Create() *errors.RestErr {

	user.DateCreated = dates.GetCurrentTimeString()

	result, err := Client.Exec(queryInsertUser, user.Firstname, user.Lastname, user.Email, user.DateCreated)

	userID, nerr := result.LastInsertId()

	if nerr != nil {
		errors.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userID

	//if currentUser != nil {
	//	if currentUser.Email == user.Email {
	//		return errors.BadRequest(fmt.Sprintf("email %s already registered", user.Email))
	//	}
	//	return errors.BadRequest(fmt.Sprintf("user %d already exists", user.Id))
	//}

	return nil
}
