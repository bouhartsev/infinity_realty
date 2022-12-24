package server

import (
	"github.com/bouhartsev/infinity_realty/internal/domain"
	"net/http"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateUserRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	response, err := s.core.CreateUser(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusCreated, w)
}
