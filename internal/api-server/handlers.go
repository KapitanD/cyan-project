package apiserver

import (
	"context"
	"encoding/json"
	"net/http"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
)

func (s *apiserver) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &v1.LoginRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		loginResponse, err := s.userService.LoginUser(context.Background(), req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, loginResponse)
	}
}

func (s *apiserver) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *apiserver) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
