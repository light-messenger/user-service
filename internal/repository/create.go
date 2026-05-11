package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) Create(ctx context.Context, nickname string, password string) (int64, error) {
	builder := sq.
		Insert(tableUsers).
		PlaceholderFormat(sq.Dollar).
		SetMap(map[string]interface{}{
			"nickname": nickname,
			"password": password,
		}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
