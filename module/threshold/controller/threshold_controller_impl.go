package controller

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/core/middleware"
	web "github.com/fathoor/posyandu-api/core/model"
	"github.com/fathoor/posyandu-api/module/threshold/model"
	"github.com/fathoor/posyandu-api/module/threshold/service"
	"github.com/gofiber/fiber/v2"
)

type thresholdControllerImpl struct {
	service.ThresholdService
}

func (controller *thresholdControllerImpl) Route(app *fiber.App) {
	threshold := app.Group("/v1/threshold", middleware.Authenticate("public"))
	threshold.Post("/", middleware.Authenticate("bidan"), controller.Create)
	threshold.Get("/", controller.GetAll)
	threshold.Get("/:parameter", controller.GetByParameter)
	threshold.Put("/:parameter", middleware.Authenticate("bidan"), controller.Update)
	threshold.Delete("/:parameter", middleware.Authenticate("bidan"), controller.Delete)
}

func (controller *thresholdControllerImpl) Create(ctx *fiber.Ctx) error {
	var request model.ThresholdCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.ThresholdService.Create(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *thresholdControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.ThresholdService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *thresholdControllerImpl) GetByParameter(ctx *fiber.Ctx) error {
	parameter := ctx.Params("parameter")

	response, err := controller.ThresholdService.GetByParameter(parameter)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *thresholdControllerImpl) Update(ctx *fiber.Ctx) error {
	parameter := ctx.Params("parameter")

	var request model.ThresholdUpdateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.ThresholdService.Update(parameter, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *thresholdControllerImpl) Delete(ctx *fiber.Ctx) error {
	parameter := ctx.Params("parameter")

	err := controller.ThresholdService.Delete(parameter)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideThresholdController(service *service.ThresholdService) ThresholdController {
	return &thresholdControllerImpl{*service}
}
