package errdomain

import "net/http"

var (
	UserNotFoundError = AppError{
		Message:  "User not found.",
		Code:     "user:not_found",
		HttpCode: http.StatusNotFound,
	}
	UserHasNeedsError = AppError{
		Message:  "User has needs dependency.",
		Code:     "user:has_needs",
		HttpCode: http.StatusBadRequest,
	}
)
