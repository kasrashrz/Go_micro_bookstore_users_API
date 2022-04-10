package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/"
	"github.com/kasrashrz/Golang_microservice/services"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"github.com/kasrashrz/Golang_microservice/utils/functionalities"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}
	ctx.JSON(http.StatusOK, result.Marshall(ctx.GetHeader("X-Public") == "true"))

}

func Read(ctx *gin.Context) {
	if oath.Ispublic != true {

	}
	userId, userErr := functionalities.GetUserID(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}

	result, err := services.UsersService.GetUser(userId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, result.Marshall(ctx.GetHeader("X-Public") == "true"))
	return
}

func Update(ctx *gin.Context) {
	var user users.User
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.BadRequestError("user id should be number")
		ctx.JSON(err.Status, err)
		return
	}

	user.Id = userId

	if err := ctx.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	isPartial := ctx.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err)
	}
	ctx.JSON(http.StatusOK, result.Marshall(ctx.GetHeader("X-Public") == "true"))

}

func Delete(ctx *gin.Context) {
	userId, userErr := functionalities.GetUserID(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}
	if err := services.UsersService.DeleteUser(userId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
	return
}

func Search(ctx *gin.Context) {
	status := ctx.Query("status")
	users, err := services.UsersService.Search(status)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, users.Marshall(ctx.GetHeader("X-Public") == "true"))
	return
}

func Login(ctx *gin.Context){
	var request users.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		restError := errors.BadRequestError("invalid json body")
		ctx.JSON(restError.Status, restError)
		return
	}
	result, saveErr := services.UsersService.LoginUser(request)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}
	ctx.JSON(http.StatusOK, result.Marshall(ctx.GetHeader("X-public") == "true"))
}