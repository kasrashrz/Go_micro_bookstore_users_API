package users

import (
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-  "`
}

// User is valid or not
func (user *User) Validate() *errors.RestErr {
	user.Firstname = strings.TrimSpace(user.Firstname)
	user.Lastname = strings.TrimSpace(user.Lastname)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequest("invalid email address")
	}
	if strings.Contains(user.Email, "@") == false {
		return errors.BadRequest("invalid email address")
	}
	return nil
}
