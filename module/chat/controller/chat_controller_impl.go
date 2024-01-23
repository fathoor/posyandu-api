package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/chat/model"
	"github.com/itsLeonB/posyandu-api/module/chat/service"
)

type chatControllerImpl struct {
	service.ChatService
}

func (controller *chatControllerImpl) Route(app *fiber.App) {
	chat := app.Group("/v1/chat", middleware.Authenticate("public"))
	chat.Post("/", controller.Create)
	chat.Post("/room", controller.CreateRoom)
	chat.Get("/room/:id", controller.GetByRoomID)
	chat.Get("/sender/:id", controller.GetBySenderID)
	chat.Put("/:id", controller.Update)
	chat.Delete("/:id", controller.Delete)
}

func (controller *chatControllerImpl) Create(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := int(claims["id"].(float64))

	var request model.ChatCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.ChatService.Create(user, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *chatControllerImpl) CreateRoom(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := int(claims["id"].(float64))

	var request model.ChatRoomCreateRequest

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.ChatService.CreateRoom(user, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *chatControllerImpl) GetByRoomID(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := int(claims["id"].(float64))

	id := ctx.Params("id")

	response, err := controller.ChatService.GetByRoomID(user, id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *chatControllerImpl) GetBySenderID(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := int(claims["id"].(float64))

	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid parameter",
		})
	}

	response, err := controller.ChatService.GetBySenderID(user, id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *chatControllerImpl) Update(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := int(claims["id"].(float64))

	var request model.ChatUpdateRequest

	id := ctx.Params("id")

	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.ChatService.Update(user, id, &request)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *chatControllerImpl) Delete(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	user := int(claims["id"].(float64))

	id := ctx.Params("id")

	err := controller.ChatService.Delete(user, id)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideChatController(service *service.ChatService) ChatController {
	return &chatControllerImpl{*service}
}
