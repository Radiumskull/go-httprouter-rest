package utils

import (
	"backend/models"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	response := models.Response{
		Message: "Success",
		Err:     nil,
		Data:    data,
	}
	response.HttpResponse(w)
}

func ErrorResponse(w http.ResponseWriter, err error) {
	response := models.Response{
		Message: err.Error(),
		Err:     err,
		Data:    nil,
	}
	response.HttpResponse(w)
}
