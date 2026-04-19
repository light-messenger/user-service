package handler

import (
	"context"
	"math"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	pb "github.com/light-messenger/user-service/pkg/userservice"
	mock "github.com/light-messenger/user-service/testutils/mocks/handler"
)

func TestGet(t *testing.T) {
	positiveId := int64(gofakeit.IntRange(1, math.MaxInt))
	negativeId := int64(gofakeit.IntRange(math.MinInt, 0))

	testCases := []struct {
		name          string
		id            int64
		validateError error
	}{
		{
			name:          "Пользователь успешно найден",
			id:            positiveId,
			validateError: nil,
		},
		{
			name:          "Передан неверный id",
			id:            negativeId,
			validateError: errIncorrectId,
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
					Get(ctx, tc.id).
					Return(gofakeit.Name(), nil)
			}

			handler := Handler{
				service: serviceMock,
			}

			_, err := handler.Get(ctx, &pb.GetRequest{
				Id: tc.id,
			})
			assert.Equal(t, tc.validateError, err, "unexpected error")
		})
	}
}
