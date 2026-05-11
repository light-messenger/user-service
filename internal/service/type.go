package service

import (
	"context"

	"github.com/light-messenger/user-service/internal/model"
)

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

type Repository interface {
	Create(ctx context.Context, nickname string, password string) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
