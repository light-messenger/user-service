package service

import "context"

func (s *Service) Get(ctx context.Context, id int64) (string, error) {
	nickname, err := s.repository.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return nickname, nil
}
