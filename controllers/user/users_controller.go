package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/services"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user);err != nil{
		restError := errors.BadRequest("invalid json body")
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
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("invalid id")
		ctx.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil{
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func UpdateUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Update\n", ctx.Param("user_id"))
}
func DeleteUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Delete\n", ctx.Param("user_id"))
}
