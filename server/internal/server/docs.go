package server

import (
	"fmt"
	"net/http"
)

func (s *Server) GetDocs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API docs")
	http.ServeFile(w, r, "docs/index.html")
}
