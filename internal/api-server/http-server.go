package apiserver

import (
	"net/http"

	"github.com/KapitanD/cyan-project/internal/config"
)

func Create(c *config.RestAPIConfig) (*http.Server, error) {
	s, err := newRESTAPIServer(c.UserServiceEndpoint)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    c.Endpoint + ":" + c.Port,
		Handler: s,
	}, nil
}
