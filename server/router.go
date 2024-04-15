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

	registerMarkRoute(mainRoute, s.db)
}

func registerMarkRoute(r fiber.Router, db *sqlx.DB) {
	ctrl := controller.NewMarkController(svc.NewMarkSvc(repo.NewMarkRepo(db)))
	markRoute := r.Group("/marks")

	newRoute(markRoute, "GET", "", ctrl.GetAllMarks)

}

func newRoute(router fiber.Router, method, path string, handler fiber.Handler) {
	router.Add(method, path, handler)
}
