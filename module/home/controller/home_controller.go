package controller

import "github.com/gofiber/fiber/v2"

type HomeController interface {
	Route(app *fiber.App)
	GetBidan(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
}
