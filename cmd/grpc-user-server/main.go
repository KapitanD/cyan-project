package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	v1 "github.com/KapitanD/cyan-project/cyan-api/proto/v1"
	user_service "github.com/KapitanD/cyan-project/pkg/user-service"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	v1.RegisterUserServiceServer(s, user_service.NewUserService(ctx))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
