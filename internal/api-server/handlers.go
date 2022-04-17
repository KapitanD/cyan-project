package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
)

func (s *apiserver) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &v1.AuthenticateRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		loginResponse, err := s.userService.AuthenticateUser(context.Background(), req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, loginResponse)
	}
}

func (s *apiserver) handleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &v1.CreateRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		createResponse, err := s.userService.CreateUser(context.Background(), req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, createResponse)
	}
}

func (s *apiserver) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		req := &v1.AuthorizeRequest{
			Token: reqToken,
		}

		authResponse, err := s.userService.AuthorizeUser(context.Background(), req)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		u := authResponse.User

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *apiserver) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*v1.UserIdentity))
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
