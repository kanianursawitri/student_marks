package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kanianursawitri/student_marks/customErr"
	"github.com/kanianursawitri/student_marks/entity"
	"github.com/kanianursawitri/student_marks/svc"
)

type RecordController struct {
	svc svc.RecordSvc
}

func NewRecordController(s svc.RecordSvc) *RecordController {
	return &RecordController{
		svc: s,
	}
}

/*
the instructions say that the request payload will be in JSON so I make the API using method POST,
room for improvement: I think it will be better if we use GET instead of POST and parse the request parameters from the URL query instead of JSON body
*/
func (c *RecordController) GetAllRecords(ctx *fiber.Ctx) error {
	var request entity.GetAllRecordsRequest
	if err := ctx.BodyParser(&request); err != nil {
		return customErr.NewBadRequestError("invalid request body")
	}

	records, err := c.svc.GetAllRecordsWithFilter(request)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"records": records,
		"code":    0,
		"msg":     "Success",
	})
}

/*
	since there's no clue how marks stored in database, I assume that we add marks through POST API
    this API only receive name and marks array, which name is the title of that marks. Make it simple, name can be anything no validation on that
*/

func (c *RecordController) AddRecord(ctx *fiber.Ctx) error {
	var request entity.AddRecordRequest
	if err := ctx.BodyParser(&request); err != nil {
		return customErr.NewBadRequestError("invalid request body")
	}

	if err := c.svc.AddRecord(request); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg":  "Record added successfully",
		"code": 0,
	})

}
