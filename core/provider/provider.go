package provider

import (
	"github.com/gofiber/fiber/v2"
	bidanController "github.com/itsLeonB/posyandu-api/module/bidan/controller"
	bidanRepository "github.com/itsLeonB/posyandu-api/module/bidan/repository"
	bidanService "github.com/itsLeonB/posyandu-api/module/bidan/service"
	posyanduController "github.com/itsLeonB/posyandu-api/module/posyandu/controller"
	posyanduRepository "github.com/itsLeonB/posyandu-api/module/posyandu/repository"
	posyanduService "github.com/itsLeonB/posyandu-api/module/posyandu/service"
	userController "github.com/itsLeonB/posyandu-api/module/user/controller"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
	userService "github.com/itsLeonB/posyandu-api/module/user/service"
	"gorm.io/gorm"
)

func ProvideModule(app *fiber.App, db *gorm.DB) {
	ProvideUser(app, db)
	ProvideBidan(app, db)
	ProvidePosyandu(app, db)
}

func ProvideUser(app *fiber.App, db *gorm.DB) {
	repository := userRepository.ProvideUserRepository(db)
	service := userService.ProvideUserService(&repository)
	controller := userController.ProvideUserController(&service)

	controller.Route(app)
}

func ProvideBidan(app *fiber.App, db *gorm.DB) {
	bidanRepo := bidanRepository.ProvideBidanRepository(db)
	userRepo := userRepository.ProvideUserRepository(db)
	service := bidanService.ProvideBidanService(&bidanRepo, &userRepo)
	controller := bidanController.ProvideBidanController(&service)

	controller.Route(app)
}

func ProvidePosyandu(app *fiber.App, db *gorm.DB) {
	repository := posyanduRepository.ProvidePosyanduRepository(db)
	service := posyanduService.ProvidePosyanduService(&repository)
	controller := posyanduController.ProvidePosyanduController(&service)

	controller.Route(app)
}
