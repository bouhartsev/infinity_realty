package server

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

// CreateProperty godoc
// @Summary CreateProperty.
// @Description
// @Tags Properties
// @Param input body domain.CreatePropertyRequest true "JSON input"
// @Success 200 {object} domain.Property
// @Failure 400 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/properties [post]
func (s *Server) CreateProperty(w http.ResponseWriter, r *http.Request) {
	var input domain.CreatePropertyRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	response, err := s.core.CreateProperty(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// GetProperty godoc
// @Summary GetProperty.
// @Description
// @Tags Properties
// @Param property_id path int true "Property id"
// @Success 200 {object} domain.Property
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/properties/{property_id} [get]
func (s *Server) GetProperty(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("property_id"))

	response, err := s.core.GetProperty(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// UpdateProperty godoc
// @Summary UpdateProperty.
// @Description
// @Tags Properties
// @Param property_id path int true "Идентификатор объекта недвижимости"
// @Param input body domain.UpdatePropertyRequest true "JSON input"
// @Success 200 {object} domain.Client
// @Failure 400 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/properties/{property_id} [patch]
func (s *Server) UpdateProperty(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdatePropertyRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	input.PropertyId, _ = strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("property_id"))

	response, err := s.core.UpdateProperty(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// DeleteProperty godoc
// @Summary DeleteProperty.
// @Description
// @Tags Properties
// @Param property_id path int true "Идентификатор объекта недвижимости"
// @Success 200
// @Failure 403 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/properties/{property_id} [delete]
func (s *Server) DeleteProperty(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("property_id"))

	err := s.core.DeleteProperty(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
