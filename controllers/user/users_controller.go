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
	if err := ctx.ShouldBindJSON(&user); err != nil {
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
		err := errors.BadRequest("user id should be number")
		ctx.JSON(err.Status, err)
		return
	}

	result, err := services.GetUser(userId)

	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
	return
}

func UpdateUser(ctx *gin.Context) {
	var user users.User

	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.BadRequest("user id should be number")
		ctx.JSON(err.Status, err)
		return
	}

	user.Id = userId

	if err := ctx.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequest("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	result, err := services.UpdateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
	}
	ctx.JSON(http.StatusOK, result)
}

func DeleteUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Delete\n", ctx.Param("user_id"))
}
