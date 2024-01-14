package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/model"
	"github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/service"
)

type jadwalPosyanduControllerImpl struct {
	service.JadwalPosyanduService
}

func (controller *jadwalPosyanduControllerImpl) Route(app *fiber.App) {
	jadwalPosyandu := app.Group("/api/jadwal-posyandu", middleware.Authenticate("bidan"))
	jadwalPosyandu.Post("/", controller.Create)
	jadwalPosyandu.Get("/", controller.GetAll)
	jadwalPosyandu.Get("/:id", controller.GetByID)
	jadwalPosyandu.Put("/:id", controller.Update)
	jadwalPosyandu.Delete("/:id", controller.Delete)
}

func (controller *jadwalPosyanduControllerImpl) Create(ctx *fiber.Ctx) error {
	var request model.JadwalPosyanduCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.JadwalPosyanduService.Create(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *jadwalPosyanduControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.JadwalPosyanduService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPosyanduControllerImpl) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.JadwalPosyanduService.GetByID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPosyanduControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.JadwalPosyanduUpdateRequest

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.JadwalPosyanduService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPosyanduControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.JadwalPosyanduService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideJadwalPosyanduController(service *service.JadwalPosyanduService) JadwalPosyanduController {
	return &jadwalPosyanduControllerImpl{*service}
}
