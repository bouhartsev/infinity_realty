package server

import (
	"net/http"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	var input domain.SignInRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	response, err := s.core.SignIn(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateUserRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	err := s.core.SignUp(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
