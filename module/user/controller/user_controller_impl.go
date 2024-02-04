package controller

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/core/middleware"
	web "github.com/fathoor/posyandu-api/core/model"
	"github.com/fathoor/posyandu-api/module/user/model"
	"github.com/fathoor/posyandu-api/module/user/service"
	"github.com/gofiber/fiber/v2"
)

type userControllerImpl struct {
	service.UserService
}

func (controller *userControllerImpl) Route(app *fiber.App) {
	auth := app.Group("/v1/auth")
	auth.Post("/login", controller.Login)

	user := app.Group("/v1/user", middleware.Authenticate("public"))
	user.Post("/register", middleware.Authenticate("kader"), controller.Register)
	user.Get("/", middleware.Authenticate("kader"), controller.GetAll)
	user.Get("/role/:role", middleware.AuthorizeRole(), controller.GetByRole)
	user.Get("/:id", middleware.AuthorizeAdminBidanOrKader(), controller.GetByID)
	user.Put("/:id", middleware.AuthorizeAdminBidanOrKader(), controller.Update)
	user.Put("/:id/auth", middleware.AuthorizeUser(), controller.UpdateAuth)
	user.Delete("/:id", middleware.Authenticate("bidan"), controller.Delete)
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
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

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
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.UserService.Update(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *userControllerImpl) UpdateAuth(ctx *fiber.Ctx) error {
	var request model.UserUpdateAuthRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.UserService.UpdateAuth(id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   "Account updated successfully!",
	})
}

func (controller *userControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	err = controller.UserService.Delete(id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideUserController(service *service.UserService) UserController {
	return &userControllerImpl{*service}
}
