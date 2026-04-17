package service

import (
	"context"

	"github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Create(ctx context.Context, nickname string, password string) (int64, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.
			WithContext(ctx).
			WithError(err).
			Error("bcrypt.GenerateFromPassword error")

		return 0, err
	}

	id, err := s.repository.Create(ctx, nickname, string(hashPassword))
	if err != nil {
		logrus.
			WithContext(ctx).
			WithField("nickname", nickname).
			WithError(err).
			Error("repository.Create error")

		return 0, err
	}

	return id, nil
}
