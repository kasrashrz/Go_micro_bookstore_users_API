package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/services"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"github.com/kasrashrz/Golang_microservice/utils/functionalities"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {
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

func Read(ctx *gin.Context) {
	userId, userErr := functionalities.GetUserID(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
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

func Update(ctx *gin.Context) {
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
	isPartial := ctx.Request.Method == http.MethodPatch
	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err)
	}
	ctx.JSON(http.StatusOK, result)
}

func Delete(ctx *gin.Context) {
	userId, userErr := functionalities.GetUserID(ctx.Param("user_id"))
	if userErr != nil {
		 ctx.JSON(userErr.Status, userErr)
		return
	}
	if err := services.DeleteUser(userId); err != nil{
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
	return
}
