package controller

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/core/middleware"
	web "github.com/fathoor/posyandu-api/core/model"
	"github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/model"
	"github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/service"
	"github.com/gofiber/fiber/v2"
)

type jadwalPenyuluhanControllerImpl struct {
	service.JadwalPenyuluhanService
}

func (controller *jadwalPenyuluhanControllerImpl) Route(app *fiber.App) {
	jadwalPenyuluhan := app.Group("/v1/jadwal-penyuluhan", middleware.Authenticate("public"))
	jadwalPenyuluhan.Post("/", middleware.Authenticate("bidan"), controller.Create)
	jadwalPenyuluhan.Get("/", controller.GetAll)
	jadwalPenyuluhan.Get("/posyandu/:id", controller.GetByPosyanduID)
	jadwalPenyuluhan.Get("/:id", controller.GetByID)
	jadwalPenyuluhan.Put("/:id", middleware.Authenticate("bidan"), controller.Update)
	jadwalPenyuluhan.Delete("/:id", middleware.Authenticate("bidan"), controller.Delete)
}

func (controller *jadwalPenyuluhanControllerImpl) Create(ctx *fiber.Ctx) error {
	var request model.JadwalPenyuluhanCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.JadwalPenyuluhanService.Create(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *jadwalPenyuluhanControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.JadwalPenyuluhanService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPenyuluhanControllerImpl) GetByPosyanduID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.JadwalPenyuluhanService.GetByPosyanduID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPenyuluhanControllerImpl) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.JadwalPenyuluhanService.GetByID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPenyuluhanControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.JadwalPenyuluhanUpdateRequest

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.JadwalPenyuluhanService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPenyuluhanControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.JadwalPenyuluhanService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideJadwalPenyuluhanController(service *service.JadwalPenyuluhanService) JadwalPenyuluhanController {
	return &jadwalPenyuluhanControllerImpl{*service}
}
