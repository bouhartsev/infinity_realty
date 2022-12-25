package errdomain

import (
	"net/http"
)

var (
	PropertyNotFoundError = AppError{
		Message:  "Property not found.",
		Code:     "property:not_found",
		HttpCode: http.StatusNotFound,
	}
	PropertyHasOffersError = AppError{
		Message:  "Property has offers.",
		Code:     "property:has_offers",
		HttpCode: http.StatusBadRequest,
	}
)
