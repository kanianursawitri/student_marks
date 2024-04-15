package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kanianursawitri/student_marks/controller"
	"github.com/kanianursawitri/student_marks/repo"
	"github.com/kanianursawitri/student_marks/svc"
)

func (s *Server) RegisterRoute() {
	mainRoute := s.app.Group("/api/v1")

	registerRecordRoute(mainRoute, s.db)
}

func registerRecordRoute(r fiber.Router, db *sqlx.DB) {
	ctrl := controller.NewRecordController(svc.NewRecordSvc(repo.NewRecordRepo(db)))
	markRoute := r.Group("/records")

	newRoute(markRoute, "POST", "/find", ctrl.GetAllRecords)
	newRoute(markRoute, "POST", "/add", ctrl.AddRecord)
}

func newRoute(router fiber.Router, method, path string, handler fiber.Handler) {
	router.Add(method, path, handler)
}
