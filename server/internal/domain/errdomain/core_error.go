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

	RealtorNotFoundError = AppError{
		Message:  "Realtor not found.",
		Code:     "realtor:not_found",
		HttpCode: http.StatusNotFound,
	}

	OfferNotFoundError = AppError{
		Message:  "Offer not found.",
		Code:     "offer:not_found",
		HttpCode: http.StatusNotFound,
	}
)
