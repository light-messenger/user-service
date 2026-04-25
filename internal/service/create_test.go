package service

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock "github.com/light-messenger/user-service/testutils/mocks/service"
)

func TestCreate(t *testing.T) {
	nickname := gofakeit.Name()
	password := gofakeit.Password(true,
		true,
		true,
		true,
		false,
		5,
	)

	testCases := []struct {
		name     string
		nickname string
		password string
	}{
		{
			name:     "Пользователь успешно создан",
			nickname: nickname,
			password: password,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ctx := context.Background()
			defer ctrl.Finish()

			repositoryMock := mock.NewMockRepository(ctrl)

			repositoryMock.
				EXPECT().
				Create(ctx, tc.nickname, gomock.All()).
				Return(gofakeit.Int64(), nil)

			service := Service{
				repository: repositoryMock,
			}

			_, err := service.Create(ctx, tc.nickname, tc.password)
			assert.Equal(t, nil, err, "unexpected error")
		})
	}
}
