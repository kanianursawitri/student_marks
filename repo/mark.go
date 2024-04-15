package repo

import "github.com/jmoiron/sqlx"

type MarkRepo interface {
}

type markRepo struct {
	db *sqlx.DB
}

func NewMarkRepo(db *sqlx.DB) MarkRepo {
	return &markRepo{
		db: db,
	}
}
