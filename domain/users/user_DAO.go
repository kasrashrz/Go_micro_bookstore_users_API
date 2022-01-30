package users

import (
	"fmt"
	. "github.com/kasrashrz/Golang_microservice/datastore/mysql/user_db"
	"github.com/kasrashrz/Golang_microservice/logger"
	"github.com/kasrashrz/Golang_microservice/utils/dates"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?, ?, ?, ?, ?, ?);"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryReadOneUser      = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryDeleteOneUser    = "DELETE FROM users WHERE id = ?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
)

func (user *User) Get() *errors.RestErr {
	statement, err := Client.Prepare(queryReadOneUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.InternalServerError("something went wrong")
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)

	if readErr := result.Scan(&user.Id,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.DateCreated,
		&user.Status); readErr != nil {

		logger.Error("error when trying to get user by id", readErr)
		return errors.InternalServerError("something went wrong")

	}
	return nil
}

func (user *User) Create() *errors.RestErr {
	statement, err := Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.InternalServerError("something went wrong")
	}

	defer statement.Close()

	user.DateCreated = dates.GetNowDbFormat()
	result, saveErr := Client.Exec(queryInsertUser,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.DateCreated,
		user.Password,
		user.Status)
	if saveErr != nil {
		logger.Error("error when trying to save user statement", saveErr)
		return errors.InternalServerError("something went wrong")
	}
	userID, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last inserted id after creating a new user", err)
		return errors.InternalServerError("something went wrong")
	}
	user.Id = userID
	return nil
}

func (user *User) Update() *errors.RestErr {
	statement, err := Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.InternalServerError("something went wrong")
	}
	defer statement.Close()

	_, err = statement.Exec(user.Firstname, user.Lastname, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.InternalServerError("something went wrong")
	}
	return nil
}

func (user *User) Delete(*User) *errors.RestErr {
	statement, err := Client.Prepare(queryDeleteOneUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.InternalServerError("something went wrong")
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Id); err != nil {
		logger.Error("error when trying to prepare delete user", err)
		return errors.InternalServerError("something went wrong")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	statement, err := Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.InternalServerError("something went wrong")
	}
	defer statement.Close()

	rows, err := statement.Query(status)

	if err != nil {
		logger.Error("error when trying to find user by status", err)
		return nil, errors.InternalServerError("something went wrong")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
			&user.DateCreated,
			&user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil,errors.InternalServerError("something went wrong")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
