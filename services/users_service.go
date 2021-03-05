package services

import (
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
)

func CreateUser(user users.User) (users.User, *errors.RestError){
	return user, nil
}
