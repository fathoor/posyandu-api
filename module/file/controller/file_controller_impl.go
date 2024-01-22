package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/middleware"
	web "github.com/itsLeonB/posyandu-api/core/model"
	"github.com/itsLeonB/posyandu-api/module/file/model"
	"github.com/itsLeonB/posyandu-api/module/file/service"
)

type fileControllerImpl struct {
	service.FileService
}

func (controller *fileControllerImpl) Route(app *fiber.App) {
	file := app.Group("/v1/file", middleware.Authenticate("public"))
	file.Post("/upload", controller.Upload)
	file.Get("/:fileType/:fileName", controller.View)
	file.Get("/:fileType/:fileName/download", controller.Download)
	file.Delete("/:fileType/:fileName", controller.Delete)
}

func (controller *fileControllerImpl) Upload(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	exception.PanicIfNeeded(err)

	fileType := ctx.FormValue("type")

	request := model.FileRequest{
		File: file,
		Type: fileType,
	}

	response, err := controller.FileService.Upload(&request)
	exception.PanicIfNeeded(err)

	err = ctx.SaveFile(file, response.URL)
	exception.PanicIfNeeded(err)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *fileControllerImpl) View(ctx *fiber.Ctx) error {
	fileType := ctx.Params("fileType")
	fileName := ctx.Params("fileName")

	response, err := controller.FileService.Get(fileType, fileName)
	exception.PanicIfNeeded(err)

	return ctx.SendFile(response)
}

func (controller *fileControllerImpl) Download(ctx *fiber.Ctx) error {
	fileType := ctx.Params("fileType")
	fileName := ctx.Params("fileName")

	response, err := controller.FileService.Get(fileType, fileName)
	exception.PanicIfNeeded(err)

	return ctx.Download(response)
}

func (controller *fileControllerImpl) Delete(ctx *fiber.Ctx) error {
	fileType := ctx.Params("fileType")
	fileName := ctx.Params("fileName")

	err := controller.FileService.Delete(fileType, fileName)
	exception.PanicIfNeeded(err)

	return ctx.SendStatus(fiber.StatusNoContent)
}

func ProvideFileController(service *service.FileService) FileController {
	return &fileControllerImpl{*service}
}
