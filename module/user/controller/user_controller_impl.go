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

	user := app.Group("/api/user", middleware.Authenticate("public"))

	user.Get("/role/:role", middleware.AuthorizeRole(), controller.GetByRole)
	user.Put("/:id", middleware.AuthorizeUser(), controller.Update)
	user.Put("/:id/auth", middleware.AuthorizeUser(), controller.UpdateAuth)
	user.Get("/:id", middleware.AuthorizeUser(), controller.GetByID)

	user.Post("/register", middleware.Authenticate("bidan"), controller.Register)
	user.Get("/", middleware.Authenticate("bidan"), controller.GetAll)
	user.Get("/:id", middleware.Authenticate("bidan"), controller.GetByID)
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
