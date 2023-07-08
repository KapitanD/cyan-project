package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	"github.com/KapitanD/cyan-project/internal/constants"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

var bearerPrefix = "Bearer "

func (s *apiserver) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &v1.AuthenticateReq{}
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
		req := &v1.UserCreateReq{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		commitId, err := s.userService.CreateUserTry(context.Background(), req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		md := metadata.Pairs(
			"email", req.Email,
		)

		_, err = s.containerService.CreateRootContainer(
			metadata.NewOutgoingContext(context.Background(), md),
			&emptypb.Empty{},
		)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			_, cancelErr := s.userService.CreateUserCancel(context.Background(), commitId)
			if cancelErr != nil {
				panic(cancelErr)
			}
			return
		}

		createResp, err := s.userService.CreateUserCommit(context.Background(), commitId)
		if err != nil {
			panic(err)
		}

		s.respond(w, r, http.StatusCreated, createResp)
	}
}

func (s *apiserver) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := s.getTokenFromReq(r)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		req := &v1.AuthorizeReq{
			Token: token,
		}

		authResponse, err := s.userService.AuthorizeUser(r.Context(), req)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		u := authResponse.User

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), constants.CtxKeyUser, u)))
	})
}

func (s *apiserver) tokenToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := s.getTokenFromReq(r)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		md := metadata.Pairs(
			string(constants.MdKeyAuthToken), token,
		)

		ctx := metadata.NewOutgoingContext(r.Context(), md)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *apiserver) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(constants.CtxKeyUser).(*v1.UserIdentity))
	}
}

func (s *apiserver) handleCreateContainer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &v1.ContainerCreateRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		resp, err := s.containerService.CreateContainer(r.Context(), req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, resp)
	}
}

func (s *apiserver) handleGetContainer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &v1.ContainerGetRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		respCont, err := s.containerService.GetContainer(r.Context(), req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		respInners, err := s.containerService.GetInners(r.Context(), req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string]interface{}{
			"container": respCont.Container,
			"inners":    respInners.Inners,
		})
	}
}

func (s *apiserver) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *apiserver) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			w.Write([]byte("can not encode response"))
		}
	}
}

func (s *apiserver) getTokenFromReq(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	n := len(bearerPrefix)

	if len(auth) < n || auth[:n] != bearerPrefix {
		return "", errors.New("auth header invalid: not bearer")
	}

	return auth[n:], nil
}
