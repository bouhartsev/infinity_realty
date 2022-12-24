package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) initRoutes() *httprouter.Router {
	r := httprouter.New()

	r.HandlerFunc(http.MethodPost, "/api/auth/sign-in", s.SignIn)
	r.HandlerFunc(http.MethodPost, "/api/auth/sign-up", s.SignUp)

	r.HandlerFunc(http.MethodGet, "/api/users/:user_id", s.GetUser)
	r.HandlerFunc(http.MethodPost, "/api/users", s.checkAuthMW(s.CreateUser))

	return r
}
