package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Go_micro_bookstore_OAth-go/oath"
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"github.com/kasrashrz/Golang_microservice/services"
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"github.com/kasrashrz/Golang_microservice/utils/functionalities"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.BadRequestError("user id should be a number")
	}
	return userId, nil
}

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
	return
}

func Read(ctx *gin.Context) {
	if err := oath.AuthenticateRequest(ctx.Request); err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	userId, idErr := getUserId(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}

	if oath.GetCallerId(ctx.Request) == user.Id {
		ctx.JSON(http.StatusOK, user.Marshall(false))
		return
	}
	ctx.JSON(http.StatusOK, user.Marshall(oath.IsPublic(ctx.Request)))
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

func Login(ctx *gin.Context) {
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
