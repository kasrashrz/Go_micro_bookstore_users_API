package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/services"
	"io/ioutil"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON();err != nil{

	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: HANDLE USER CREATION ERROR
		return
	}
	ctx.JSON(http.StatusCreated, result)
}
func ReadUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Read\n")
}
func UpdateUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Update\n", ctx.Param("user_id"))
}
func DeleteUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Delete\n", ctx.Param("user_id"))
}
