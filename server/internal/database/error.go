package database

import "errors"

var (
	ErrUserHasNeeds      = errors.New("user has needs")
	ErrPropertyHasOffers = errors.New("property has offers")
)
