package api

import (
	"net/http"
	"portfolio-api/usecases"
)

type ErrorResponseObject struct {
	Message string `json:message`
}

func GetErrorResponse(err *usecases.UError) (int, ErrorResponseObject) {
	return getErrorHTTPStatus(err.Code), ErrorResponseObject{
		Message: err.Msg,
	}
}

func getErrorHTTPStatus(errCode int) int {
	switch errCode {
	case usecases.BadRequest:
		return http.StatusBadRequest
	case usecases.Unauthorized:
		return http.StatusUnauthorized
	case usecases.NotFound:
		return http.StatusNotFound
	case usecases.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}