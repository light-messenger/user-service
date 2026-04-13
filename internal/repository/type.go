package repository

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const (
	tableUsers = "users"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
