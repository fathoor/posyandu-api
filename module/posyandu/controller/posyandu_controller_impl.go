package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/posyandu/model"
	"github.com/itsLeonB/posyandu-api/module/posyandu/service"
)

type posyanduControllerImpl struct {
	service.PosyanduService
}

func (controller *posyanduControllerImpl) Route(app *fiber.App) {
	posyandu := app.Group("/api/posyandu", middleware.Authenticate("bidan"))
	posyandu.Post("/", controller.Create)
	posyandu.Get("/", controller.GetAll)
	posyandu.Get("/:id", controller.GetByID)
	posyandu.Put("/:id", controller.Update)
	posyandu.Delete("/:id", controller.Delete)
}

func (controller *posyanduControllerImpl) Create(ctx *fiber.Ctx) error {
	var request model.PosyanduCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.PosyanduService.Create(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *posyanduControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.PosyanduService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *posyanduControllerImpl) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.PosyanduService.GetByID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *posyanduControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.PosyanduUpdateRequest

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.PosyanduService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *posyanduControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.PosyanduService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvidePosyanduController(service *service.PosyanduService) PosyanduController {
	return &posyanduControllerImpl{*service}
}
