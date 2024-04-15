package svc

import (
	"github.com/kanianursawitri/student_marks/customErr"
	"github.com/kanianursawitri/student_marks/entity"
	"github.com/kanianursawitri/student_marks/repo"
)

type RecordSvc interface {
	AddRecord(request entity.AddRecordRequest) error
}

type recordSvc struct {
	repo repo.RecordRepo
}

func NewRecordSvc(r repo.RecordRepo) RecordSvc {
	return &recordSvc{
		repo: r,
	}
}

func (s *recordSvc) AddRecord(request entity.AddRecordRequest) error {
	mark := entity.NewRecord(request.Name, request.Marks)

	if err := mark.Validate(); err != nil {
		return customErr.NewBadRequestError(err.Error())
	}

	if err := s.repo.AddRecord(request); err != nil {
		return err
	}

	return nil
}
