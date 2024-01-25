package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/pemeriksaan/model"
	"github.com/itsLeonB/posyandu-api/module/pemeriksaan/service"
)

type pemeriksaanControllerImpl struct {
	service.PemeriksaanService
}

func (controller *pemeriksaanControllerImpl) Route(app *fiber.App) {
	pemeriksaan := app.Group("/v1/pemeriksaan", middleware.Authenticate("public"))
	pemeriksaan.Post("/", middleware.Authenticate("bidan"), controller.Create)
	pemeriksaan.Get("/", middleware.Authenticate("bidan"), controller.GetAll)
	pemeriksaan.Get("/remaja/:id", middleware.AuthorizeAdminOrBidan(), controller.GetByRemajaUserID)
	pemeriksaan.Get("/:id", controller.GetByID)
	pemeriksaan.Put("/:id", middleware.Authenticate("bidan"), controller.Update)
	pemeriksaan.Delete("/:id", middleware.Authenticate("bidan"), controller.Delete)
}

func (controller *pemeriksaanControllerImpl) Create(ctx *fiber.Ctx) error {
	var request model.PemeriksaanCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.PemeriksaanService.Create(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *pemeriksaanControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.PemeriksaanService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pemeriksaanControllerImpl) GetByRemajaUserID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.PemeriksaanService.GetByRemajaUserID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pemeriksaanControllerImpl) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.PemeriksaanService.GetByID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pemeriksaanControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.PemeriksaanUpdateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.PemeriksaanService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *pemeriksaanControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.PemeriksaanService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvidePemeriksaanController(service *service.PemeriksaanService) PemeriksaanController {
	return &pemeriksaanControllerImpl{*service}
}
