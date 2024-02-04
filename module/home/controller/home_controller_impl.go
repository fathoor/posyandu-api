package controller

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/core/middleware"
	web "github.com/fathoor/posyandu-api/core/model"
	"github.com/fathoor/posyandu-api/module/home/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type homeControllerImpl struct {
	service.HomeService
}

func (controller *homeControllerImpl) Route(app *fiber.App) {
	home := app.Group("/v1/home", middleware.Authenticate("public"))
	home.Get("/bidan", middleware.Authenticate("bidan"), controller.GetBidan)
	home.Get("/", controller.Get)
}

func (controller *homeControllerImpl) GetBidan(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))

	response, err := controller.HomeService.GetBidan(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *homeControllerImpl) Get(ctx *fiber.Ctx) error {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))

	response, err := controller.HomeService.Get(id)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func ProvideHomeController(service *service.HomeService) HomeController {
	return &homeControllerImpl{*service}
}
