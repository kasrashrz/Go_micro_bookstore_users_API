package services

import (
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := users.User.Validate(&user); err != nil{
		return  nil, err
	}
	return nil, nil
}
