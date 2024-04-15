package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	"github.com/kanianursawitri/student_marks/middleware"
)

type Server struct {
	db        *sqlx.DB
	app       *fiber.App
	validator *validator.Validate
}

func NewServer(db *sqlx.DB) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	validate := validator.New()

	app.Use(recover.New())

	return &Server{
		db:        db,
		app:       app,
		validator: validate,
	}
}

func (s *Server) Run() error {
	return s.app.Listen(":8080")
}
