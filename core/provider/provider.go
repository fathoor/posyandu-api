package provider

import (
	"github.com/gofiber/fiber/v2"
	bidanController "github.com/itsLeonB/posyandu-api/module/bidan/controller"
	bidanRepository "github.com/itsLeonB/posyandu-api/module/bidan/repository"
	bidanService "github.com/itsLeonB/posyandu-api/module/bidan/service"
	pengampuController "github.com/itsLeonB/posyandu-api/module/pengampu/controller"
	pengampuRepository "github.com/itsLeonB/posyandu-api/module/pengampu/repository"
	pengampuService "github.com/itsLeonB/posyandu-api/module/pengampu/service"
	posyanduController "github.com/itsLeonB/posyandu-api/module/posyandu/controller"
	posyanduRepository "github.com/itsLeonB/posyandu-api/module/posyandu/repository"
	posyanduService "github.com/itsLeonB/posyandu-api/module/posyandu/service"
	remajaController "github.com/itsLeonB/posyandu-api/module/remaja/controller"
	remajaRepository "github.com/itsLeonB/posyandu-api/module/remaja/repository"
	remajaService "github.com/itsLeonB/posyandu-api/module/remaja/service"
	userController "github.com/itsLeonB/posyandu-api/module/user/controller"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
	userService "github.com/itsLeonB/posyandu-api/module/user/service"
	"gorm.io/gorm"
)

func ProvideModule(app *fiber.App, db *gorm.DB) {
	ProvideUser(app, db)
	ProvideBidan(app, db)
	ProvidePosyandu(app, db)
	ProvideRemaja(app, db)
	ProvidePengampu(app, db)
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

func ProvideRemaja(app *fiber.App, db *gorm.DB) {
	remajaRepo := remajaRepository.ProvideRemajaRepository(db)
	posyanduRepo := posyanduRepository.ProvidePosyanduRepository(db)
	userRepo := userRepository.ProvideUserRepository(db)
	service := remajaService.ProvideRemajaService(&remajaRepo, &posyanduRepo, &userRepo)
	controller := remajaController.ProvideRemajaController(&service)

	controller.Route(app)
}

func ProvidePengampu(app *fiber.App, db *gorm.DB) {
	pengampuRepo := pengampuRepository.ProvidePengampuRepository(db)
	bidanRepo := bidanRepository.ProvideBidanRepository(db)
	posyanduRepo := posyanduRepository.ProvidePosyanduRepository(db)
	userRepo := userRepository.ProvideUserRepository(db)
	service := pengampuService.ProvidePengampuService(&bidanRepo, &pengampuRepo, &posyanduRepo, &userRepo)
	controller := pengampuController.ProvidePengampuController(&service)

	controller.Route(app)
}
