package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/pengampu/model"
	"github.com/itsLeonB/posyandu-api/module/pengampu/service"
)

type pengampuControllerImpl struct {
	service.PengampuService
}

func (controller *pengampuControllerImpl) Route(app *fiber.App) {
	pengampu := app.Group("/api/pengampu", middleware.Authenticate("bidan"))
	pengampu.Post("/", controller.Create)
	pengampu.Get("/", controller.GetAll)
	pengampu.Get("/:id", controller.GetByID)
	pengampu.Put("/:id", controller.Update)
	pengampu.Delete("/:id", controller.Delete)
}

func (controller *pengampuControllerImpl) Create(ctx *fiber.Ctx) error {
	var request model.PengampuCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.PengampuService.Create(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *pengampuControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.PengampuService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pengampuControllerImpl) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.PengampuService.GetByID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pengampuControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.PengampuUpdateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.PengampuService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pengampuControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.PengampuService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvidePengampuController(service *service.PengampuService) PengampuController {
	return &pengampuControllerImpl{*service}
}
