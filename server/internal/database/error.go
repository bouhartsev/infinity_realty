package database

import "errors"

var (
	ErrClientHasNeeds    = errors.New("client has needs")
	ErrPropertyHasOffers = errors.New("property has offers")
)
