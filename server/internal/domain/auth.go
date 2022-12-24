package domain

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	jwt.RegisteredClaims
	User User `json:"user"`
}
