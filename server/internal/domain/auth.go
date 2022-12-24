package domain

import "github.com/golang-jwt/jwt/v4"

type (
	AuthClaims struct {
		jwt.RegisteredClaims
		User User `json:"user"`
	}

	SignInRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	SignInResponse struct {
		Token string `json:"token"`
		User  User   `json:"user"`
	}
)
