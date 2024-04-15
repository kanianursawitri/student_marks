package svc

import (
	"github.com/kanianursawitri/student_marks/customErr"
	"github.com/kanianursawitri/student_marks/entity"
	"github.com/kanianursawitri/student_marks/repo"
)

type RecordSvc interface {
	AddRecord(request entity.AddRecordRequest) error
	GetAllRecordsWithFilter(filter entity.GetAllRecordsRequest) ([]entity.RecordResponse, error)
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

	if err := s.repo.AddRecord(mark); err != nil {
		return err
	}

	return nil
}

func (s *recordSvc) GetAllRecordsWithFilter(request entity.GetAllRecordsRequest) ([]entity.RecordResponse, error) {
	filter, err := entity.NewFilter(request)
	if err != nil {
		return nil, customErr.NewBadRequestError(err.Error())
	}

	if err := filter.Validate(); err != nil {
		return nil, customErr.NewBadRequestError(err.Error())
	}

	records, err := s.repo.GetAllRecordsWithFilter(filter)
	if err != nil {
		return nil, err
	}

	responses := make([]entity.RecordResponse, len(records))
	for i, record := range records {
		responses[i] = entity.RecordResponse{
			ID:         record.ID,
			TotalMarks: record.TotalMarks,
			CreatedAt:  record.CreatedAt.Format("2006-01-02T15:04:05.000Z"),
		}
	}

	return responses, nil
}
