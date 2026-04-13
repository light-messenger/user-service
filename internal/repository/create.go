package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) Create(ctx context.Context, nickname string, password string) (int64, error) {
	builder := sq.
		Insert(tableUsers).
		SetMap(map[string]interface{}{
			"nickname": nickname,
			"password": password,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
