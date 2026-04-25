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

func TestGet(t *testing.T) {
	id := gofakeit.Int64()
	nickname := gofakeit.Name()

	testCases := []struct {
		name   string
		query  string
		args   []driver.Value
		rows   *sqlmock.Rows
		result string
		err    error
	}{
		{
			name:   "Успешное получение пользователя",
			query:  `SELECT nickname FROM users WHERE id = ?`,
			args:   []driver.Value{id},
			rows:   sqlmock.NewRows([]string{"nickname"}).AddRow(nickname),
			result: nickname,
			err:    nil,
		},
		{
			name:  "Ошибка при получении пользователя",
			query: `SELECT nickname FROM users WHERE id = ?`,
			args:  []driver.Value{id},
			rows:  sqlmock.NewRows([]string{"nickname"}),
			err:   sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, dbMock, _ := sqlmock.New()
			ctx := context.Background()
			defer db.Close()

			dbMock.
				ExpectQuery(regexp.QuoteMeta(tc.query)).
				WithArgs(tc.args...).
				WillReturnRows(tc.rows).
				WillReturnError(tc.err)

			repository := Repository{
				db: db,
			}

			nickname, err := repository.Get(ctx, id)
			assert.Equal(t, tc.result, nickname, "unexpected result")
			assert.Equal(t, tc.err, err, "unexpected error")
		})
	}
}
