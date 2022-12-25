package errdomain

import "net/http"

var (
	RealtorNotFoundError = AppError{
		Message:  "Realtor not found.",
		Code:     "realtor:not_found",
		HttpCode: http.StatusNotFound,
	}
)
