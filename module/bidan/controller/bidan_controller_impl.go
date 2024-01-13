package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/bidan/model"
	"github.com/itsLeonB/posyandu-api/module/bidan/service"
)

type bidanControllerImpl struct {
	service.BidanService
}

func (controller *bidanControllerImpl) Route(app *fiber.App) {
	bidan := app.Group("/api/bidan", middleware.Authenticate("bidan"))
	bidan.Post("/", controller.Create)
	bidan.Get("/", controller.GetAll)
	bidan.Get("/:id", controller.GetByID)
	bidan.Put("/:id", controller.Update)
	bidan.Delete("/:id", controller.Delete)
}

func (controller *bidanControllerImpl) Create(ctx *fiber.Ctx) error {
	var request model.BidanCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.BidanService.Create(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *bidanControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.BidanService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *bidanControllerImpl) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.BidanService.GetByID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *bidanControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.BidanUpdateRequest

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.BidanService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *bidanControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.BidanService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideBidanController(service *service.BidanService) BidanController {
	return &bidanControllerImpl{*service}
}
