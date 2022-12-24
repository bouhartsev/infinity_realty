package errdomain

var (
	InvalidAuthTokenError = AppError{
		Message: "Invalid auth token.",
		Code:    "auth:invalid_token",
	}
	ForbiddenError = AppError{
		Message: "You have no access for this resource.",
		Code:    "auth:forbidden",
	}
)
