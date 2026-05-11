package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/light-messenger/user-service/internal/model"
)

func (r *Repository) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.
		Select("id", "nickname", "registered_at").
		From(tableUsers).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user model.User
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
