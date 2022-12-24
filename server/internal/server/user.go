package server

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

// CreateUser godoc
// @Summary Создает пользователя.
// @Description
// @Tags Users
// @Param input body domain.CreateUserRequest true "JSON input"
// @Success 200 {object} domain.User
// @Failure 400 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/users [post]
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

// GetUser godoc
// @Summary Возвращает информацию о пользователе.
// @Description
// @Tags Users
// @Param user_id path int true "Идентификатор пользователя"
// @Success 200 {object} domain.User
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/users/{user_id} [get]
func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("user_id"))

	response, err := s.core.GetUser(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// UpdateUser godoc
// @Summary Обновляет пользователя.
// @Description
// @Tags Users
// @Param user_id path int true "Идентификатор пользователя"
// @Param input body domain.UpdateUserRequest true "JSON input"
// @Success 200 {object} domain.User
// @Failure 400 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/users/{user_id} [post]
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

// DeleteUser godoc
// @Summary Удаляет пользователя.
// @Description
// @Tags Users
// @Param user_id path int true "Идентификатор пользователя"
// @Success 200
// @Failure 403 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/users/{user_id} [get]
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("user_id"))

	err := s.core.DeleteUser(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
