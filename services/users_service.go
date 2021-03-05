package services

import "github.com/kasrashrz/Golang_microservice/domain/users"

func CreateUser(user users.User) (*users.User, error){

	return &user, nil
}