package server

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/bouhartsev/infinity_realty/internal/domain"
)

// CreateOffer godoc
// @Summary CreateOffer.
// @Description
// @Tags Offers
// @Param input body domain.CreateOfferRequest true "JSON input"
// @Success 200 {object} domain.Offer
// @Failure 400 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/offers [post]
func (s *Server) CreateOffer(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateOfferRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	response, err := s.core.CreateOffer(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// GetOffer godoc
// @Summary GetOffer.
// @Description
// @Tags Offers
// @Param offer_id path int true "Идентификатор клиента"
// @Success 200 {object} domain.Offer
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/offers/{offer_id} [get]
func (s *Server) GetOffer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("offer_id"))

	response, err := s.core.GetOffer(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// UpdateOffer godoc
// @Summary UpdateOffer.
// @Description
// @Tags Offers
// @Param offer_id path int true "Идентификатор клиента"
// @Param input body domain.UpdateOfferRequest true "JSON input"
// @Success 200 {object} domain.Offer
// @Failure 400 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/offers/{offer_id} [patch]
func (s *Server) UpdateOffer(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateOfferRequest
	if err := s.ReadJson(&input, r, w); err != nil {
		return
	}

	input.OfferId, _ = strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("offer_id"))

	response, err := s.core.UpdateOffer(r.Context(), &input)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	s.Json(response, http.StatusOK, w)
}

// DeleteOffer godoc
// @Summary DeleteOffer.
// @Description
// @Tags Offers
// @Param offer_id path int true "Идентификатор клиента"
// @Success 200
// @Failure 403 {object} errdomain.AppError
// @Failure 404 {object} errdomain.AppError
// @Failure 500 {object} errdomain.AppError
// @Router /api/offers/{offer_id} [delete]
func (s *Server) DeleteOffer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("offer_id"))

	err := s.core.DeleteOffer(r.Context(), id)
	if err != nil {
		s.ParseError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
