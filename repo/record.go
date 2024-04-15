package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanianursawitri/student_marks/entity"
)

type RecordRepo interface {
	AddRecord(request *entity.Record) error
	GetAllRecordsWithFilter(filter *entity.GetAllRecordsFilter) ([]entity.RecordResult, error)
}

type recordRepo struct {
	db *sqlx.DB
}

func NewRecordRepo(db *sqlx.DB) RecordRepo {
	return &recordRepo{
		db: db,
	}
}

func (r *recordRepo) AddRecord(request *entity.Record) error {
	statement := "INSERT INTO records (name, marks) VALUES ($1, $2)"
	stmt, err := r.db.Prepare(statement)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(request.Name, request.Marks)
	return err
}

func (r *recordRepo) GetAllRecordsWithFilter(filter *entity.GetAllRecordsFilter) ([]entity.RecordResult, error) {
	var records []entity.RecordResult
	query :=
		`SELECT id, 
				(SELECT SUM(elem::int) FROM unnest(marks) AS t(elem)) AS total_marks,
				created_at 
 		FROM records
 		WHERE created_at BETWEEN $1 AND $2
   			  AND (SELECT SUM(elem::int) FROM unnest(marks) AS t(elem)) BETWEEN $3 AND $4`

	err := r.db.Select(&records, query, filter.StartDate, filter.EndDate, filter.MinCount, filter.MaxCount)
	if err != nil {
		return nil, err
	}

	return records, nil
}
