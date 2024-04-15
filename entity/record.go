package entity

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AddRecordRequest struct {
	Name  string `json:"name"`
	Marks []int  `json:"marks"`
}

/*
 example of get all request payload
 {
	"startDate": "2016-01-26",
	"endDate": "2018-02-02",
	"minCount": 100,
	"maxCount": 300
 }
*/

type GetAllRecordsRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type GetAllRecordsFilter struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	MinCount  int       `json:"minCount"`
	MaxCount  int       `json:"maxCount"`
}

type Record struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Marks     []int     `db:"marks"`
	CreatedAt time.Time `db:"created_at"`
}

type RecordResult struct {
	ID         int       `db:"id"`
	TotalMarks int       `db:"total_marks"`
	CreatedAt  time.Time `db:"created_at"`
}

type RecordResponse struct {
	ID         int    `json:"id"`
	TotalMarks int    `json:"totalMarks"`
	CreatedAt  string `json:"createdAt"`
}

func NewRecord(name string, marks []int) *Record {
	return &Record{
		Name:  name,
		Marks: marks,
	}
}

func NewFilter(request GetAllRecordsRequest) (*GetAllRecordsFilter, error) {
	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		return nil, errors.New("invalid startDate format")
	}

	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		return nil, errors.New("invalid endDate format")
	}

	filter := &GetAllRecordsFilter{
		StartDate: startDate,
		EndDate:   endDate,
		MinCount:  request.MinCount,
		MaxCount:  request.MaxCount,
	}
	return filter, nil
}

func (m *Record) Validate() error {
	err := validation.ValidateStruct(m,
		validation.Field(&m.Name,
			validation.Required.Error("name is required"),
			validation.Length(1, 255).Error("name must be 1-255 characters"),
		),
		validation.Field(&m.Marks,
			validation.Required.Error("marks is required"),
			validation.Length(1, 0).Error("marks must contain at least 1 item"),
		),
	)

	return err
}

func (f *GetAllRecordsFilter) Validate() error {
	err := validation.ValidateStruct(f,
		validation.Field(&f.StartDate,
			validation.Required.Error("startDate is required"),
		),
		validation.Field(&f.EndDate,
			validation.Required.Error("endDate is required"),
		),
		validation.Field(&f.MinCount,
			validation.Required.Error("minCount is required"),
		),
		validation.Field(&f.MaxCount,
			validation.Required.Error("maxCount is required"),
		),
	)

	if f.EndDate.Before(f.StartDate) {
		return errors.New("endDate must be after startDate")
	}

	return err
}
