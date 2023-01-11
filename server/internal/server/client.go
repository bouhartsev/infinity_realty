package server

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

// CreateClient godoc
// @Summary CreateClient.
// @Description
// @Tags Clients
// @Param input body domain.CreateClientRequest true "Также обязательны либо `email`, либо `telephone`, либо оба."
// @Success 200 {object} domain.Client
// @Failure 400 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/clients [post]
func (s *Server) CreateClient(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateClientRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	response, err := s.core.CreateClient(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// GetClient godoc
// @Summary GetClient.
// @Description
// @Tags Clients
// @Param client_id path int true "Идентификатор клиента"
// @Success 200 {object} domain.Client
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/clients/{client_id} [get]
func (s *Server) GetClient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("client_id"))

	response, err := s.core.GetClient(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// UpdateClient godoc
// @Summary UpdateClient.
// @Description
// @Tags Clients
// @Param client_id path int true "Идентификатор клиента"
// @Param input body domain.UpdateClientRequest true "JSON input"
// @Success 200 {object} domain.Client
// @Failure 400 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/clients/{client_id} [patch]
func (s *Server) UpdateClient(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateClientRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	input.ClientId, _ = strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("client_id"))

	response, err := s.core.UpdateClient(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// DeleteClient godoc
// @Summary DeleteClient.
// @Description
// @Tags Clients
// @Param client_id path int true "Идентификатор клиента"
// @Success 200
// @Failure 403 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/clients/{client_id} [delete]
func (s *Server) DeleteClient(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("client_id"))

	err := s.core.DeleteClient(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
