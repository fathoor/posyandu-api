package controller

import "github.com/gofiber/fiber/v2"

type ChatController interface {
	Route(app *fiber.App)
	Create(ctx *fiber.Ctx) error
	CreateRoom(ctx *fiber.Ctx) error
	GetByRoomID(ctx *fiber.Ctx) error
	GetBySenderID(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
