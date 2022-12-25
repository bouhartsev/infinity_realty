package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) initRoutes() *httprouter.Router {
	r := httprouter.New()

	r.HandlerFunc(http.MethodPost, "/api/auth/sign-in", s.SignIn)
	r.HandlerFunc(http.MethodPost, "/api/auth/sign-up", s.SignUp)

	r.HandlerFunc(http.MethodPost, "/api/users", s.checkAuthMW(s.CreateUser))
	r.HandlerFunc(http.MethodGet, "/api/users/:user_id", s.checkAuthMW(s.GetUser))
	r.HandlerFunc(http.MethodPatch, "/api/users/:user_id", s.checkAuthMW(s.UpdateUser))
	r.HandlerFunc(http.MethodDelete, "/api/users/:user_id", s.checkAuthMW(s.DeleteUser))

	r.HandlerFunc(http.MethodPost, "/api/properties", s.checkAuthMW(s.CreateProperty))
	r.HandlerFunc(http.MethodGet, "/api/properties/:property_id", s.checkAuthMW(s.GetProperty))
	r.HandlerFunc(http.MethodPatch, "/api/properties/:property_id", s.checkAuthMW(s.UpdateProperty))
	r.HandlerFunc(http.MethodDelete, "/api/properties/:property_id", s.checkAuthMW(s.DeleteProperty))

	r.HandlerFunc(http.MethodGet, "/api/docs", s.GetDocs)

	return r
}
