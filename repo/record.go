package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanianursawitri/student_marks/entity"
)

type RecordRepo interface {
	AddRecord(request entity.AddRecordRequest) error
}

type recordRepo struct {
	db *sqlx.DB
}

func NewRecordRepo(db *sqlx.DB) RecordRepo {
	return &recordRepo{
		db: db,
	}
}

func (r *recordRepo) AddRecord(request entity.AddRecordRequest) error {
	statement := "INSERT INTO records (name, marks) VALUES ($1, $2)"
	stmt, err := r.db.Prepare(statement)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(request.Name, request.Marks)
	return err
}
