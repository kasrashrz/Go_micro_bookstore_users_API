package main

import (
	"github.com/kasrashrz/Golang_microservice/app"
	"github.com/kasrashrz/Golang_microservice/logger"
)

func main() {
	logger.Info("the application is about to start")
	app.StartApplication()
}
