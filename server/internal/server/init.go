package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) initRoutes() *httprouter.Router {
	r := httprouter.New()

	r.HandlerFunc(http.MethodPost, "/api/clients", s.CreateClient)
	r.HandlerFunc(http.MethodGet, "/api/clients/:client_id", s.GetClient)
	r.HandlerFunc(http.MethodPatch, "/api/clients/:client_id", s.UpdateClient)
	r.HandlerFunc(http.MethodDelete, "/api/clients/:client_id", s.DeleteClient)

	r.HandlerFunc(http.MethodPost, "/api/realtors", s.CreateRealtor)
	r.HandlerFunc(http.MethodGet, "/api/realtors/:realtor_id", s.GetRealtor)
	r.HandlerFunc(http.MethodPatch, "/api/realtors/:realtor_id", s.UpdateRealtor)
	r.HandlerFunc(http.MethodDelete, "/api/realtors/:realtor_id", s.DeleteRealtor)

	r.HandlerFunc(http.MethodPost, "/api/offers", s.CreateOffer)
	r.HandlerFunc(http.MethodGet, "/api/offers/:offer_id", s.GetOffer)
	r.HandlerFunc(http.MethodPatch, "/api/offers/:offer_id", s.UpdateOffer)
	r.HandlerFunc(http.MethodDelete, "/api/offers/:offer_id", s.DeleteOffer)

	r.HandlerFunc(http.MethodPost, "/api/properties", s.CreateProperty)
	r.HandlerFunc(http.MethodGet, "/api/properties/:property_id", s.GetProperty)
	r.HandlerFunc(http.MethodPatch, "/api/properties/:property_id", s.UpdateProperty)
	r.HandlerFunc(http.MethodDelete, "/api/properties/:property_id", s.DeleteProperty)

	r.HandlerFunc(http.MethodGet, "/api/docs", s.GetDocs)

	return r
}
