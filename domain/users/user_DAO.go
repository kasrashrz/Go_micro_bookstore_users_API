package users

import (
	"fmt"
	. "github.com/kasrashrz/Golang_microservice/datastore/mysql/user_db"
	"github.com/kasrashrz/Golang_microservice/utils/dates"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "Duplicate entry"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryReadOneUser = "SELECT * FROM users WHERE id = ?;"
)

func (user *User) Get() *errors.RestErr {
	statement, err := Client.Prepare(queryReadOneUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email , &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows){
			return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.InternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Create() *errors.RestErr {
	statement, err := Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	defer statement.Close()

	user.DateCreated = dates.GetCurrentTimeString()

	result, err := Client.Exec(queryInsertUser, user.Firstname, user.Lastname, user.Email, user.DateCreated)

	if err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.BadRequest(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.BadRequest(fmt.Sprintf("user %d already exists", user.Id))
	}

	userID, nerr := result.LastInsertId()

	if nerr != nil {
		errors.InternalServerError(fmt.Sprintf("error when trying to save user: %v", err.Error()))
	}

	user.Id = userID

	return nil
}
