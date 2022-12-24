package server

import "github.com/julienschmidt/httprouter"

func (s *Server) initRoutes() *httprouter.Router {
	r := httprouter.New()

	return r
}
