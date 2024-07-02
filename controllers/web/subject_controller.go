package web

import (
	"github.com/gofiber/fiber/v2"
	"go_starter/controllers"
	"go_starter/requests"
	"go_starter/services"
	"go_starter/validation"
)

type SubjectController interface {
	CreateSubjectService(ctx *fiber.Ctx) error
	FilterSubjectBySubjectCodeService(ctx *fiber.Ctx) error
	GetSubjectController(ctx *fiber.Ctx) error
}

type subjectController struct {
	serviceSubjectService services.SubjectService
}

func (s *subjectController) CreateSubjectService(ctx *fiber.Ctx) error {
	request := new(requests.InsertSubjectRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := s.serviceSubjectService.CreateSubjectService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response.Message)
}

func (s *subjectController) FilterSubjectBySubjectCodeService(ctx *fiber.Ctx) error {
	request := new(requests.SubjectRequest)
	if err := ctx.BodyParser(request); err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	errValidate := validation.Validate(request)
	if errValidate != nil {
		return controllers.NewErrorValidate(ctx, errValidate[0].Error)
	}
	response, err := s.serviceSubjectService.FilterSubjectBySubjectCodeService(*request)
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)

}

func (s *subjectController) GetSubjectController(ctx *fiber.Ctx) error {
	response, err := s.serviceSubjectService.GetSubjectService()
	if err != nil {
		return controllers.NewErrorResponses(ctx, err)
	}
	return controllers.NewSuccessResponse(ctx, response)
}

func NewSubjectController(
	serviceSubjectService services.SubjectService,
) SubjectController {
	return &subjectController{
		serviceSubjectService: serviceSubjectService,
	}
}
