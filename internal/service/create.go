package service

import (
	"context"
)

func (s *Service) Create(ctx context.Context, nickname string, password string) (int64, error) {
	id, err := s.repository.Create(ctx, nickname, password)
	if err != nil {
		return 0, err
	}

	return id, nil
}
