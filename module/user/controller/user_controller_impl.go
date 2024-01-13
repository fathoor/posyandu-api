package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/user/model"
	"github.com/itsLeonB/posyandu-api/module/user/service"
)

type userControllerImpl struct {
	service.UserService
}

func (controller *userControllerImpl) Route(app *fiber.App) {
	auth := app.Group("/api/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/forget-password", controller.ForgetPassword)

	bidan := app.Group("/api/user", middleware.Authenticate("bidan"))
	bidan.Post("/register", controller.Register)
	bidan.Get("/", controller.GetAll)
	bidan.Put("/:id", controller.Update)
	bidan.Delete("/:id", controller.Delete)

	public := app.Group("/api/user", middleware.Authenticate("public"))
	public.Get("/role/:role", controller.GetByRole)
	public.Get("/:id", controller.GetByID)
}

func (controller *userControllerImpl) Login(ctx *fiber.Ctx) error {
	var request model.LoginRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.UserService.Login(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *userControllerImpl) ForgetPassword(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (controller *userControllerImpl) Register(ctx *fiber.Ctx) error {
	var request model.UserRegisterRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.UserService.Register(&request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *userControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := controller.UserService.GetAll()
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *userControllerImpl) GetByRole(ctx *fiber.Ctx) error {
	role := ctx.Params("role")

	response, err := controller.UserService.GetByRole(role)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *userControllerImpl) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	exception.PanicIfNeeded(err)

	response, err := controller.UserService.GetByID(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *userControllerImpl) Update(ctx *fiber.Ctx) error {
	var request model.UserUpdateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	id, err := ctx.ParamsInt("id")
	exception.PanicIfNeeded(err)

	response, err := controller.UserService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *userControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	exception.PanicIfNeeded(err)

	err = controller.UserService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideUserController(service *service.UserService) UserController {
	return &userControllerImpl{*service}
}
