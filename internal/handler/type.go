package handler

import (
	"context"
	"errors"

	pb "github.com/light-messenger/user-service/pkg/userservice"
)

var (
	errEmptyNickname = errors.New("empty nickname")
	errEmptyPassword = errors.New("empty password")
	errIncorrectId   = errors.New("incorrect id")
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
