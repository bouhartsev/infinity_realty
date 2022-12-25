package server

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

// CreateRealtor godoc
// @Summary CreateRealtor.
// @Description
// @Tags Realtors
// @Param input body domain.CreateRealtorRequest true "JSON input"
// @Success 200 {object} domain.Realtor
// @Failure 400 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/realtors [post]
func (s *Server) CreateRealtor(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateRealtorRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	response, err := s.core.CreateRealtor(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// GetRealtor godoc
// @Summary GetRealtor.
// @Description
// @Tags Realtors
// @Param realtor_id path int true "Идентификатор клиента"
// @Success 200 {object} domain.Realtor
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/realtors/{realtor_id} [get]
func (s *Server) GetRealtor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("realtor_id"))

	response, err := s.core.GetRealtor(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// UpdateRealtor godoc
// @Summary UpdateRealtor.
// @Description
// @Tags Realtors
// @Param realtor_id path int true "Идентификатор клиента"
// @Param input body domain.UpdateRealtorRequest true "JSON input"
// @Success 200 {object} domain.Client
// @Failure 400 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/realtors/{realtor_id} [patch]
func (s *Server) UpdateRealtor(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateRealtorRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	input.RealtorId, _ = strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("realtor_id"))

	response, err := s.core.UpdateRealtor(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// DeleteRealtor godoc
// @Summary DeleteRealtor.
// @Description
// @Tags Realtors
// @Param realtor_id path int true "Идентификатор клиента"
// @Success 200
// @Failure 403 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/realtors/{realtor_id} [delete]
func (s *Server) DeleteRealtor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("realtor_id"))

	err := s.core.DeleteRealtor(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
