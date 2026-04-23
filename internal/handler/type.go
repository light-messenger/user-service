package handler

import (
	"context"

	pb "github.com/light-messenger/user-service/pkg/userservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errEmptyNickname = status.Error(codes.InvalidArgument, "empty nickname")
	errEmptyPassword = status.Error(codes.InvalidArgument, "empty password")
	errIncorrectId   = status.Error(codes.InvalidArgument, "incorrect id")
)

type Handler struct {
	pb.UnimplementedUserServiceServer

	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

type Service interface {
	Create(ctx context.Context, nickname string, password string) (int64, error)
	Get(ctx context.Context, id int64) (string, error)
}
