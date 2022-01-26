package users

import (
	. "github.com/kasrashrz/Golang_microservice/datastore/mysql/user_db"
	"github.com/kasrashrz/Golang_microservice/utils/dates"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"github.com/kasrashrz/Golang_microservice/utils/mysql_utils"
)

const (
	queryInsertUser    = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryUpdateUser    = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryReadOneUser   = "SELECT * FROM users WHERE id = ?;"
	queryDeleteOneUser = "DELETE FROM users WHERE id = ?;"
)

func (user *User) Get() *errors.RestErr {
	statement, err := Client.Prepare(queryReadOneUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)

	if readErr := result.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated); readErr != nil {
		return mysql_utils.ParseError(readErr)
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

	result, saveErr := Client.Exec(queryInsertUser, user.Firstname, user.Lastname, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userID
	return nil
}

func (user *User) Update() *errors.RestErr {
	statement, err := Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	_, err = statement.Exec(user.Firstname, user.Lastname, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete(*User) *errors.RestErr {
	statement, err := Client.Prepare(queryDeleteOneUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}
