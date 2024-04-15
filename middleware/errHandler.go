package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kanianursawitri/student_marks/customErr"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := ctx.Response().StatusCode()
	if customError, ok := err.(customErr.CustomError); ok {
		code = customError.Status()
		return ctx.Status(code).JSON(fiber.Map{
			"msg":  customError.Error(),
			"code": code,
		})
	} else if code < 400 {
		code = fiber.StatusInternalServerError
	}

	return ctx.Status(code).JSON(fiber.Map{
		"msg":  err.Error(),
		"code": code,
	})
}
