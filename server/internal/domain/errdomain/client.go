package errdomain

import "net/http"

var (
	ClientNotFoundError = AppError{
		Message:  "Client not found.",
		Code:     "client:not_found",
		HttpCode: http.StatusNotFound,
	}
	ClientHasNeedsError = AppError{
		Message:  "Client has needs dependency.",
		Code:     "client:has_needs",
		HttpCode: http.StatusBadRequest,
	}
)
