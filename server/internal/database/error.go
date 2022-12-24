package database

import "errors"

var (
	ErrUserHasNeeds = errors.New("user has needs")
	ErrUserHasDeals = errors.New("user has deals")
)
