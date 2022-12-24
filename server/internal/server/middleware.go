package server

import (
	"context"
	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func (s *Server) checkAuthMW(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path

		if p == `/api/auth/sign-in` || strings.HasPrefix(p, `/api/docs`) {
			handler(w, r)
			return
		}

		header := r.Header.Get("Authorization")

		if header == "" {
			s.Json(errdomain.InvalidAuthTokenError, http.StatusUnauthorized, w)
			return
		}

		vals := strings.Split(header, "Bearer ")

		if len(vals) != 2 {
			s.Json(errdomain.InvalidAuthTokenError, http.StatusUnauthorized, w)
			return
		}

		accessToken := vals[1]

		claims := &domain.AuthClaims{}

		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.cfg.TokenKey), nil
		})

		if _, ok := token.Claims.(*domain.AuthClaims); !ok {
			s.Json(errdomain.InvalidAuthTokenError, http.StatusUnauthorized, w)
			return
		}

		if err != nil {
			s.Json(errdomain.InvalidAuthTokenError, http.StatusUnauthorized, w)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "user", claims.User))
		handler(w, r)
	}
}
