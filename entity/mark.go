package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AddRecordRequest struct {
	Name  string `json:"name"`
	Marks []int  `json:"marks"`
}

type Record struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Marks     []int     `db:"marks"`
	CreatedAt time.Time `db:"created_at"`
}

func NewRecord(name string, marks []int) *Record {
	return &Record{
		Name:  name,
		Marks: marks,
	}
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
