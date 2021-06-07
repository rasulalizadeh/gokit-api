package router

import "net/http"

type Response struct {
	Status  uint        `json:"status"`
	Message string      `json:"message"`
	Errors  []Error    `json:"errors"`
	Data    interface{} `json:"data"`
}

func NewResponse(status uint, message string, data interface{}, errors ...Error) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}

func NewSuccessResponse(message string, data interface{}) *Response {
	return NewResponse(http.StatusOK, message, data)
}
func NewErrorResponse(message string, error Error) *Response {
	return NewResponse(error.Code, message, nil, error)
}
func NewNotFoundResponse(message string, errorData interface{}) *Response {
	return NewErrorResponse(message, NewModelNotFoundError(message, errorData))
}
func NewForbiddenResponse(message string, errorData interface{}) *Response {
	return NewErrorResponse(message, NewForbiddenError(message, errorData))
}
func NewBadRequestResponse(message string, errorData interface{}) *Response {
	return NewErrorResponse(message, NewBadRequestError(message, errorData))
}
func NewValidationFailResponse(message string, errorData interface{}) *Response {
	return NewErrorResponse(message, NewValidationError(message, errorData))
}