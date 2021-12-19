package apiserver

import (
	"net/http"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type apiserver struct {
	router      *mux.Router
	userService v1.UserServiceClient
}

func newRESTAPIServer(userServerEndpoint string) (*apiserver, error) {
	conn, err := grpc.Dial(userServerEndpoint)
	if err != nil {
		return nil, err
	}

	srv := &apiserver{
		router:      mux.NewRouter(),
		userService: v1.NewUserServiceClient(conn),
	}

	srv.configureRouter()

	return srv, nil
}

func (s *apiserver) configureRouter() {
	s.router.HandleFunc("/login", s.handleLogin()).Methods("GET")
}

func (s *apiserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
