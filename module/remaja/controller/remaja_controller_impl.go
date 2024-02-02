package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/remaja/model"
	"github.com/itsLeonB/posyandu-api/module/remaja/service"
)

type remajaControllerImpl struct {
	service.RemajaService
}

func (controller *remajaControllerImpl) Route(app *fiber.App) {
	remaja := app.Group("/v1/remaja")
	remaja.Post("/", middleware.Authenticate("bidan"), controller.Create)
	remaja.Get("/", middleware.Authenticate("bidan"), controller.GetAll)
	remaja.Get("/posyandu/:id", middleware.Authenticate("public"), controller.GetByPosyanduID)
	remaja.Get("/:id", middleware.Authenticate("bidan"), controller.GetByID)
	remaja.Put("/:id", middleware.Authenticate("bidan"), controller.UpdateKader)
	remaja.Delete("/:id", middleware.Authenticate("bidan"), controller.Delete)

	kader := app.Group("/v1/kader")
	kader.Get("/", middleware.Authenticate("public"), controller.GetAllKader)
	kader.Post("/remaja", middleware.Authenticate("kader"), controller.Create)
	kader.Get("/remaja", middleware.Authenticate("kader"), controller.GetAll)
	kader.Get("/remaja/:id", middleware.Authenticate("kader"), controller.GetByID)
	kader.Put("/remaja/:id", middleware.Authenticate("kader"), controller.Update)
}

func (controller *remajaControllerImpl) Create(ctx *fiber.Ctx) error {
	var request model.RemajaCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.RemajaService.Create(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *remajaControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.RemajaService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *remajaControllerImpl) GetAllKader(ctx *fiber.Ctx) error {
	response, err := controller.RemajaService.GetAllKader()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *remajaControllerImpl) GetByPosyanduID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.RemajaService.GetByPosyanduID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *remajaControllerImpl) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.RemajaService.GetByID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *remajaControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.RemajaUpdateRequest

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.RemajaService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *remajaControllerImpl) UpdateKader(ctx *fiber.Ctx) error {
	var request model.RemajaUpdateKaderRequest

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.RemajaService.UpdateKader(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *remajaControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.RemajaService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideRemajaController(service *service.RemajaService) RemajaController {
	return &remajaControllerImpl{*service}
}
