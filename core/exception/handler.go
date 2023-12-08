package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itsLeonB/posyandu-api/core/model"
)

func Handler(ctx *fiber.Ctx, err error) error {
	switch err.(type) {
	case BadRequestError:
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	case UnauthorizedError:
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Code:   fiber.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   err.Error(),
		})
	case ForbiddenError:
		return ctx.Status(fiber.StatusForbidden).JSON(model.Response{
			Code:   fiber.StatusForbidden,
			Status: "Forbidden",
			Data:   err.Error(),
		})
	case NotFoundError:
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   fiber.StatusNotFound,
			Status: "Not Found",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
		Code:   fiber.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err.Error(),
	})
}
