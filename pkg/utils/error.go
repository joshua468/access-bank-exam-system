package utils

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (e *AppError) Error() string {
	return e.Message
}

func HandleError(err error) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return NewAppError(http.StatusInternalServerError, "Internal Server Error")
}

func ResponseError(err error) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: err.Error()}
}

func ResponseSuccess(message string) *AppError {
	return &AppError{Code: http.StatusOK, Message: message}
}
