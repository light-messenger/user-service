package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	nickname := gofakeit.Name()
	password := gofakeit.Password(
		true,
		true,
		true,
		true,
		false,
		5,
	)

	testCases := []struct {
		name   string
		query  string
		args   []driver.Value
		result sql.Result
		want   int64
		err    error
	}{
		{
			name:   "Успешное создание пользователя",
			query:  `INSERT INTO users (nickname,password) VALUES (?,?)`,
			args:   []driver.Value{nickname, password},
			result: sqlmock.NewResult(1, 1),
			want:   1,
			err:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, dbMock, _ := sqlmock.New()
			ctx := context.Background()
			defer db.Close()

			dbMock.
				ExpectExec(regexp.QuoteMeta(tc.query)).
				WithArgs(tc.args...).
				WillReturnResult(tc.result)

			repository := Repository{
				db: db,
			}

			id, err := repository.Create(ctx, nickname, password)
			assert.Equal(t, tc.want, id, "unexpected result")
			assert.Equal(t, tc.err, err, "unexpected error")
		})
	}
}
