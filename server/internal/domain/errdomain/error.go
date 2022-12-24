package errdomain

import "net/http"

type AppError struct {
	Message  string  `json:"message"`
	Code     string  `json:"code"`
	HttpCode int     `json:"-"`
	Field    *string `json:"field,omitempty"`
}

func (e AppError) Error() string {
	return e.Message
}

func NewInternalError(msg string) AppError {
	return AppError{
		Message:  msg,
		HttpCode: http.StatusInternalServerError,
		Code:     "internal",
	}
}

func NewInvalidRequestError(msg string) AppError {
	return AppError{
		Message:  msg,
		HttpCode: http.StatusBadRequest,
		Code:     "invalid_request",
	}
}

func NewUserError(msg, field string) AppError {
	return AppError{
		Message:  msg,
		HttpCode: http.StatusBadRequest,
		Code:     "invalid_input",
		Field:    &field,
	}
}
