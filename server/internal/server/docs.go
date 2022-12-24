package server

import "net/http"

func (s *Server) GetDocs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "docs/index.html")
}
