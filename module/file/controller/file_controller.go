package controller

import "github.com/gofiber/fiber/v2"

type FileController interface {
	Route(app *fiber.App)
	Upload(ctx *fiber.Ctx) error
	View(ctx *fiber.Ctx) error
	Download(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
