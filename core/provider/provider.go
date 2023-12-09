package provider

import (
	"github.com/gofiber/fiber/v2"
	userController "github.com/itsLeonB/posyandu-api/module/user/controller"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
	userService "github.com/itsLeonB/posyandu-api/module/user/service"
	"gorm.io/gorm"
)

func ProvideModule(app *fiber.App, db *gorm.DB) {
	ProvideUser(app, db)
}

func ProvideUser(app *fiber.App, db *gorm.DB) {
	repository := userRepository.ProvideUserRepository(db)
	service := userService.ProvideUserService(&repository)
	controller := userController.ProvideUserController(&service)

	controller.Route(app)
}
