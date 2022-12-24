package server

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/bouhartsev/infinity_realty/internal/domain"
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

	s.Json(response, http.StatusOK, w)
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("user_id"))

	response, err := s.core.GetUser(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateUserRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	input.UserId, _ = strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("user_id"))

	response, err := s.core.UpdateUser(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("user_id"))

	err := s.core.DeleteUser(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
