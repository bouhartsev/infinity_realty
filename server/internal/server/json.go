package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

func (s *Server) ReadJson(data any, r *http.Request, w http.ResponseWriter) error {
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		s.Json(errdomain.NewInternalError(err.Error()), http.StatusInternalServerError, w)
		return err
	}
	err = json.Unmarshal(buf, data)
	if err != nil {
		s.Json(errdomain.NewInternalError(err.Error()), http.StatusInternalServerError, w)
		return err
	}
	return nil
}

func (s *Server) Json(data any, status int, w http.ResponseWriter) {
	buf, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(buf)
}
