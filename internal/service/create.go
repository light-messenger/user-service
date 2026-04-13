package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (s *Service) Create(ctx context.Context, nickname string, password string) (int64, error) {
	id, err := s.repository.Create(ctx, nickname, password)
	if err != nil {
		logrus.
			WithContext(ctx).
			WithFields(logrus.Fields{
				"nickname": nickname,
				"password": password,
			}).
			WithError(err).
			Error("repository.Create error")

		return 0, err
	}

	return id, nil
}
