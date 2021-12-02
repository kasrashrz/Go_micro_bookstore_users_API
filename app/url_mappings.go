package app

import (
	"github.com/kasrashrz/Golang_microservice/controllers/ping"
	"github.com/kasrashrz/Golang_microservice/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/read/:user_id", user.ReadUser)
	router.POST("/users/create", user.CreateUser)
	router.PUT("/users/update/:user_id", user.UpdateUser)
	router.DELETE("/users/delete/:user_id", user.DeleteUser)
}
