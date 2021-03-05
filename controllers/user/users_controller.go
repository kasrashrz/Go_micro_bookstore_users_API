package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/services"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user);err != nil{
		restError := errors.NewBadRequest("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
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
