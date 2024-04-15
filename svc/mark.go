package svc

import "github.com/kanianursawitri/student_marks/repo"

type MarkSvc interface {
}

type markSvc struct {
	repo repo.MarkRepo
}

func NewMarkSvc(r repo.MarkRepo) MarkSvc {
	return &markSvc{
		repo: r,
	}
}
