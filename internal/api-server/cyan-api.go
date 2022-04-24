package apiserver

import (
	"net/http"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type apiserver struct {
	router           *mux.Router
	userService      v1.UserServiceClient
	containerService v1.ContainerServiceClient
}

func newRESTAPIServer(userServerEndpoint, containerServerEndpoint string) (*apiserver, error) {
	userConn, err := grpc.Dial(userServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	containerConn, err := grpc.Dial(containerServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	srv := &apiserver{
		router:           mux.NewRouter(),
		userService:      v1.NewUserServiceClient(userConn),
		containerService: v1.NewContainerServiceClient(containerConn),
	}

	srv.configureRouter()

	return srv, nil
}

func (s *apiserver) configureRouter() {
	s.router.HandleFunc("/login", s.handleLogin()).Methods("POST")
	s.router.HandleFunc("/register", s.handleUserCreate()).Methods("POST")
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
	containers := s.router.PathPrefix("/containers").Subrouter()
	containers.Use(s.tokenToContext)
	containers.HandleFunc("/", s.handleCreateContainer()).Methods("POST")
	containers.HandleFunc("/", s.handleGetContainer()).Methods("GET")
}

func (s *apiserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
