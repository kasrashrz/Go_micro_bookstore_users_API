package services

import (
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/utils/crypto_utils"
	"github.com/kasrashrz/Golang_microservice/utils/dates"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dates.GetNowDbFormat()
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Create(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if !isPartial {
		if user.Firstname != "" {
			current.Firstname = user.Firstname
		}
		if user.Lastname != "" {
			current.Lastname = user.Lastname
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.Firstname = user.Firstname
		current.Lastname = user.Lastname
		current.Email = user.Email
	}

	if err := user.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userID int64) *errors.RestErr {
	current, err := GetUser(userID)
	if err != nil {
		return err
	}
	if err := current.Delete(current); err != nil {
		return err
	}
	return nil
}

func Search(status string) (users.Users, *errors.RestErr) {
	dao := users.User{}
	return dao.FindByStatus(status)
}
