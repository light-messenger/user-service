package service

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock "github.com/light-messenger/user-service/testutils/mocks/service"
)

func TestGet(t *testing.T) {
	id := gofakeit.Int64()

	testCases := []struct {
		name string
		id   int64
	}{
		{
			name: "Пользователь успешно найден",
			id:   id,
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
				Get(ctx, tc.id).
				Return(gofakeit.Name(), nil)

			service := Service{
				repository: repositoryMock,
			}

			_, err := service.Get(ctx, tc.id)
			assert.Equal(t, nil, err, "unexpected error")
		})
	}
}
