package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	apiserver "github.com/KapitanD/cyan-project/internal/api-server"
	"github.com/KapitanD/cyan-project/internal/config"
)

func main() {
	config := &config.RestAPIConfig{
		Endpoint:            "localhost",
		Port:                "8080",
		UserServiceEndpoint: "localhost:10051",
	}

	server, err := apiserver.Create(config)
	if err != nil {
		log.Fatal().Err(err).Msgf("cant create server")
	}

	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Err(err).Msgf("HTTP server Shutdown")
		}
		log.Info().Msg("server gracefully shutdown")
		close(idleConsClosed)
	}()

	log.Info().Msgf("starting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatal().Err(err).Msg("HTTP server ListenAndServe")
	}
}
