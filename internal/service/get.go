package service

import (
	"context"

	"github.com/light-messenger/user-service/internal/model"
	"github.com/sirupsen/logrus"
)

func (s *Service) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.repository.Get(ctx, id)
	if err != nil {
		logrus.
			WithContext(ctx).
			WithField("id", id).
			WithError(err).
			Error("repository.Get error")

		return nil, err
	}

	return user, nil
}
