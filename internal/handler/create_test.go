package handler

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	pb "github.com/light-messenger/user-service/pkg/userservice"
	mock "github.com/light-messenger/user-service/testutils/mocks/handler"
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
		name          string
		nickname      string
		password      string
		validateError error
	}{
		{
			name:          "Пользователь успешно создан",
			nickname:      nickname,
			password:      password,
			validateError: nil,
		},
		{
			name:          "Не передан nickname",
			nickname:      "",
			password:      password,
			validateError: errEmptyNickname,
		},
		{
			name:          "Не передан password",
			nickname:      nickname,
			password:      "",
			validateError: errEmptyPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ctx := context.Background()
			defer ctrl.Finish()

			serviceMock := mock.NewMockService(ctrl)

			if tc.validateError == nil {
				serviceMock.
					EXPECT().
					Create(ctx, tc.nickname, tc.password).
					Return(gofakeit.Int64(), nil)
			}

			handler := Handler{
				service: serviceMock,
			}

			_, err := handler.Create(ctx, &pb.CreateRequest{
				Nickname: tc.nickname,
				Password: tc.password,
			})
			assert.Equal(t, tc.validateError, err, "unexpected error")
		})
	}
}
