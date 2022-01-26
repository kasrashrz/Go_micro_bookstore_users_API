package functionalities

import (
	"github.com/kasrashrz/Golang_microservice/utils/errors"
	"strconv"
)

func GetUserID(givenUserID string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(givenUserID, 10, 64)
	if userErr != nil {
		return 0, errors.BadRequest("user id should be number")
	}
	return userId, nil

}
