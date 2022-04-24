package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	"github.com/KapitanD/cyan-project/internal/config"
	container_service "github.com/KapitanD/cyan-project/pkg/container-service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	config := &config.ContainerServiceConfig{
		UserServiceEndpoint: "localhost:10051",
	}
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	s := grpc.NewServer()

	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

		<-sigint

		// We received an interrupt signal, shut down.
		s.GracefulStop()

		log.Info().Msg("server gracefully shutdown")
		close(idleConsClosed)
	}()

	containerService, err := container_service.NewContainerService(config)
	if err != nil {
		log.Fatal().Err(err).Msg("fail to create container service")
	}

	v1.RegisterContainerServiceServer(s, containerService)

	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
