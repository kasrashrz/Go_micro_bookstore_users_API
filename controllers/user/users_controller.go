package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/services"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user);err != nil{
		fmt.Println(err)
		//TODO: RETURN BAD REQUEST TO CALLER
		return
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
