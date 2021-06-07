package router

import "net/http"

type Error struct {
	Code    uint        `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewModelNotFoundError(message string, data interface{}) Error {
	return Error{
		Code:    http.StatusNotFound,
		Type:    "ModelNotFoundError",
		Message: message,
		Data:    data,
	}
}
func NewForbiddenError(message string, data interface{}) Error {
	return Error{
		Code:    http.StatusForbidden,
		Type:    "ForbiddenError",
		Message: message,
		Data:    data,
	}
}
func NewBadRequestError(message string, data interface{}) Error {
	return Error{
		Code:    http.StatusBadRequest,
		Type:    "BadRequestError",
		Message: message,
		Data:    data,
	}
}
func NewValidationError(message string, data interface{}) Error {
	return Error{
		Code:    http.StatusUnprocessableEntity,
		Type:    "ValidationError",
		Message: message,
		Data:    data,
	}
}
