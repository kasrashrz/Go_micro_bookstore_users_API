package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequest(message string)*RestError{
	return &RestError{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}