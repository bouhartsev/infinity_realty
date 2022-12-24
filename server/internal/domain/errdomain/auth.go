package errdomain

import "net/http"

var (
	InvalidAuthTokenError = AppError{
		Message:  "Invalid auth token.",
		HttpCode: http.StatusUnauthorized,
		Code:     "auth:invalid_token",
	}
	ForbiddenError = AppError{
		Message:  "You have no access for this resource.",
		HttpCode: http.StatusForbidden,
		Code:     "auth:forbidden",
	}
	InvalidCredentialsError = AppError{
		Message:  "Invalid credentials.",
		Code:     "invalid_credentials",
		HttpCode: http.StatusBadRequest,
	}
)
