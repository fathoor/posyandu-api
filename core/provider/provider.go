package provider

import (
	"github.com/gofiber/fiber/v2"
	bidanController "github.com/itsLeonB/posyandu-api/module/bidan/controller"
	bidanRepository "github.com/itsLeonB/posyandu-api/module/bidan/repository"
	bidanService "github.com/itsLeonB/posyandu-api/module/bidan/service"
	fileController "github.com/itsLeonB/posyandu-api/module/file/controller"
	fileService "github.com/itsLeonB/posyandu-api/module/file/service"
	jadwalPenyuluhanController "github.com/itsLeonB/posyandu-api/module/jadwal-penyuluhan/controller"
	jadwalPenyuluhanRepository "github.com/itsLeonB/posyandu-api/module/jadwal-penyuluhan/repository"
	jadwalPenyuluhanService "github.com/itsLeonB/posyandu-api/module/jadwal-penyuluhan/service"
	jadwalPosyanduController "github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/controller"
	jadwalPosyanduRepository "github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/repository"
	jadwalPosyanduService "github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/service"
	pemeriksaanController "github.com/itsLeonB/posyandu-api/module/pemeriksaan/controller"
	pemeriksaanRepository "github.com/itsLeonB/posyandu-api/module/pemeriksaan/repository"
	pemeriksaanService "github.com/itsLeonB/posyandu-api/module/pemeriksaan/service"
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
	ProvideJadwalPosyandu(app, db)
	ProvideJadwalPenyuluhan(app, db)
	ProvidePemeriksaan(app, db)
	ProvideFile(app)
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

func ProvideJadwalPosyandu(app *fiber.App, db *gorm.DB) {
	jadwalPosyanduRepo := jadwalPosyanduRepository.ProvideJadwalPosyanduRepository(db)
	posyanduRepo := posyanduRepository.ProvidePosyanduRepository(db)
	service := jadwalPosyanduService.ProvideJadwalPosyanduService(&jadwalPosyanduRepo, &posyanduRepo)
	controller := jadwalPosyanduController.ProvideJadwalPosyanduController(&service)

	controller.Route(app)
}

func ProvideJadwalPenyuluhan(app *fiber.App, db *gorm.DB) {
	jadwalPenyuluhanRepo := jadwalPenyuluhanRepository.ProvideJadwalPenyuluhanRepository(db)
	posyanduRepo := posyanduRepository.ProvidePosyanduRepository(db)
	service := jadwalPenyuluhanService.ProvideJadwalPenyuluhanService(&jadwalPenyuluhanRepo, &posyanduRepo)
	controller := jadwalPenyuluhanController.ProvideJadwalPenyuluhanController(&service)

	controller.Route(app)
}

func ProvidePemeriksaan(app *fiber.App, db *gorm.DB) {
	pemeriksaanRepo := pemeriksaanRepository.ProvidePemeriksaanRepository(db)
	posyanduRepo := posyanduRepository.ProvidePosyanduRepository(db)
	remajaRepo := remajaRepository.ProvideRemajaRepository(db)
	userRepo := userRepository.ProvideUserRepository(db)
	service := pemeriksaanService.ProvidePemeriksaanService(&pemeriksaanRepo, &posyanduRepo, &remajaRepo, &userRepo)
	controller := pemeriksaanController.ProvidePemeriksaanController(&service)

	controller.Route(app)
}

func ProvideFile(app *fiber.App) {
	service := fileService.ProvideFileService()
	controller := fileController.ProvideFileController(&service)

	controller.Route(app)
}
