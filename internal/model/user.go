package model

import "time"

type User struct {
	Id           int64     `db:"id"`
	Nickname     string    `db:"nickname"`
	RegisteredAt time.Time `db:"registered_at"`
}
