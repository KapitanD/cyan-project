PROJECT?=github.com/KapitanD/cyan-project

.PHONY: build
protogen:
	protoc --proto_path=. --go_out=. --go-grpc_out=. cyan-api/api/v1/*.proto

.PHONY: build
build: protogen
	CGO_ENABLED=0 go build -o ./build/rest-api-server -v ./cmd/rest-api-server
	CGO_ENABLED=0 go build -o ./build/grpc-user-server -v ./cmd/grpc-user-server
	CGO_ENABLED=0 go build -o ./build/grpc-container-server -v ./cmd/grpc-container-server

.DEFAULT_GOAL := build
