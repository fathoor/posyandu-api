package controller

import "github.com/gofiber/fiber/v2"

type RemajaController interface {
	Route(app *fiber.App)
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetAllKader(ctx *fiber.Ctx) error
	GetByPosyanduID(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	UpdateKader(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
