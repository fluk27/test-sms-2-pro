package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}
func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
func NewInternalServerError(message string) error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
func NewBadRequest(message string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}
func NewConflict(message string) error {
	return AppError{
		Code:    http.StatusConflict,
		Message: message,
	}
}
