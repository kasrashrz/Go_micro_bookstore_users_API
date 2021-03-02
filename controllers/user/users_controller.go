package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kasrashrz/Golang_microservice/domain/users"
	"io/ioutil"
	"net/http"
)

//type UserController interface {
//	CreateUser()
//	ReadUser()
//	UpdateUser()
//	DeleteUser()
//}

func CreateUser(ctx *gin.Context) {
	var user users.User
	fmt.Print(user)
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println(err)
	fmt.Println(string(bytes))
	ctx.String(http.StatusOK, "Create\n")
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
