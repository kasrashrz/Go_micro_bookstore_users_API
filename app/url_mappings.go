package app

import (
	"github.com/kasrashrz/Golang_microservice/controllers/ping"
	"github.com/kasrashrz/Golang_microservice/controllers/user"
)

func mapUrls () {
	router.GET("/", ping.Ping)
	router.POST("/users/create/:user_id", user.CreateUser)
	router.POST("/users/read", user.ReadUser)
	router.POST("/users/update/:user_id", user.UpdateUser)
	router.POST("/users/delete/:user_id", user.DeleteUser)
}