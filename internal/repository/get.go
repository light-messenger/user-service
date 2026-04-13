package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) Get(ctx context.Context, id int64) (string, error) {
	builder := sq.
		Select("nickname").
		From(tableUsers).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	var nickname string
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&nickname)
	if err != nil {
		return "", err
	}

	return nickname, nil
}
