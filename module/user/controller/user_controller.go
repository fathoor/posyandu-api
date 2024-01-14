package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Route(app *fiber.App)
	Login(ctx *fiber.Ctx) error
	ForgetPassword(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByRole(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	UpdateAuth(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
