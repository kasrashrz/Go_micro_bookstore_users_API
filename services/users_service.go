package services

import (
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Create(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser (userId int64) (*users.User, *errors.RestErr){
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}