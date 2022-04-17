package apiserver

import (
	"net/http"

	"github.com/KapitanD/cyan-project/internal/config"
)

func Create(conf *config.RestAPIConfig) (*http.Server, error) {
	s, err := newRESTAPIServer(conf.UserServiceEndpoint)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    conf.Endpoint + ":" + conf.Port,
		Handler: s,
	}, nil
}
