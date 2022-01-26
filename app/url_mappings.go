package app

import (
	"github.com/kasrashrz/Golang_microservice/controllers/ping"
	"github.com/kasrashrz/Golang_microservice/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.Create)
	router.GET("/users/:user_id", user.Read)
	router.PUT("/users/:user_id", user.Update)
	router.PATCH("/users/:user_id", user.Update)
	router.DELETE("/users/:user_id", user.Delete)
}
