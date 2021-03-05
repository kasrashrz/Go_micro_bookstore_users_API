package services

import "github.com/kasrashrz/Golang_microservice/domain/users"

func CreateUser(user users.User) (*users.User, error){

	return &user, nil
}

//obj = {
//"1":{
//  "message": "User XXX Not Found",
//  "status": 404,
//  "error": "Not Found"
//},
//{
//  ""
//}
//}