package services

import (
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"strings"
)

func CreateUser(user users.User) (users.User, *errors.RestError) {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	return user, nil
}
