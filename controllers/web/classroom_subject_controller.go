package web

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go_starter/controllers"
	"go_starter/requests"
	"go_starter/services"
	"go_starter/validation"
)

type ClassroomSubjectController interface {
	GetClassroomSubjectController(ctx *fiber.Ctx) error
	CreateClassroomSubjectController(ctx *fiber.Ctx) error
}

type classroomSubjectController struct {
	serviceClassroomSubject services.ClassroomSubjectService
}

func (c *classroomSubjectController) GetClassroomSubjectController(ctx *fiber.Ctx) error {
	response, err := c.serviceClassroomSubject.GetClassroomSubjectService()
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func (c *classroomSubjectController) CreateClassroomSubjectController(ctx *fiber.Ctx) error {
	request := new(requests.CreateClassroomSubjectRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	fmt.Println(request)
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := c.serviceClassroomSubject.CreateClassroomSubjectService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}

	return controllers.NewSuccessMessage(ctx, response.Message)
}

func NewClassroomSubjectController(
	serviceClassroomSubject services.ClassroomSubjectService,
) ClassroomSubjectController {
	return &classroomSubjectController{
		serviceClassroomSubject: serviceClassroomSubject,
	}
}
