package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kanianursawitri/student_marks/svc"
)

type MarkController struct {
	svc svc.MarkSvc
}

func NewMarkController(s svc.MarkSvc) *MarkController {
	return &MarkController{
		svc: s,
	}
}

func (c *MarkController) GetAllMarks(ctx *fiber.Ctx) error {
	return nil
}
