package main

import (
	// "context"
	"net"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	handler "github.com/light-messenger/user-service/internal/handler"
	repository "github.com/light-messenger/user-service/internal/repository"
	service "github.com/light-messenger/user-service/internal/service"
	userService "github.com/light-messenger/user-service/pkg/userservice"
	"github.com/sirupsen/logrus"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("USER_SERVICE_DSN")
	if dsn == "" {
		logrus.Fatal("USER_SERVICE_DSN is required")
	}

	db, err := sqlx.Open(
		"pgx",
		dsn,
	)
	if err != nil {
		logrus.
			WithError(err).
			Fatal("sql.Open fatal error")
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
		logrus.
			WithField("address", address).
			WithError(err).
			Fatal("net.Listen error")
	}

	logrus.Infof("gRPC server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		logrus.
			WithError(err).
			Fatal("grpcServer.Serve error")
	}
}
