package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *Server) initRoutes() *httprouter.Router {
	r := httprouter.New()

	r.HandlerFunc(http.MethodPost, "/api/users", s.checkAuthMW(s.CreateUser))

	return r
}
