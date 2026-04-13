package main

import (
	// "context"
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "modernc.org/sqlite"

	handler "github.com/light-messenger/user-service/internal/handler"
	repository "github.com/light-messenger/user-service/internal/repository"
	service "github.com/light-messenger/user-service/internal/service"
	userService "github.com/light-messenger/user-service/pkg/userservice"
)

func main() {
	db, err := sql.Open("sqlite", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repository.New(db)
	service := service.New(repository)
	handler := handler.New(service)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	userService.RegisterUserServiceServer(grpcServer, handler)

	address := "localhost:6666"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	log.Printf("gRPC server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
