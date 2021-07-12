package utils

import (
	"backend/models"
	"net/http"
)

func SuccessResponseWithMessage(w http.ResponseWriter, data interface{}, message string) {
	response := models.Response{
		Message: message,
		Err:     nil,
		Data:    data,
	}
	response.HttpResponse(w)
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	SuccessResponseWithMessage(w, data, "Success")
}

func ErrorResponseWithMessage(w http.ResponseWriter, err error, message string) {
	response := models.Response{
		Message: err.Error(),
		Err:     err,
		Data:    nil,
	}
	response.HttpResponse(w)
}

func ErrorResponse(w http.ResponseWriter, err error) {
	ErrorResponseWithMessage(w, err, err.Error())
}
