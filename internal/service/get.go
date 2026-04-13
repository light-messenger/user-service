package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (s *Service) Get(ctx context.Context, id int64) (string, error) {
	nickname, err := s.repository.Get(ctx, id)
	if err != nil {
		logrus.
			WithContext(ctx).
			WithField("id", id).
			WithError(err).
			Error("repository.Get error")

		return "", err
	}

	return nickname, nil
}
