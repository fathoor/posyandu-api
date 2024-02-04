package provider

import (
	bidanController "github.com/fathoor/posyandu-api/module/bidan/controller"
	bidanRepository "github.com/fathoor/posyandu-api/module/bidan/repository"
	bidanService "github.com/fathoor/posyandu-api/module/bidan/service"
	chatController "github.com/fathoor/posyandu-api/module/chat/controller"
	chatRepository "github.com/fathoor/posyandu-api/module/chat/repository"
	chatService "github.com/fathoor/posyandu-api/module/chat/service"
	fileController "github.com/fathoor/posyandu-api/module/file/controller"
	fileService "github.com/fathoor/posyandu-api/module/file/service"
	homeController "github.com/fathoor/posyandu-api/module/home/controller"
	homeService "github.com/fathoor/posyandu-api/module/home/service"
	jadwalPenyuluhanController "github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/controller"
	jadwalPenyuluhanRepository "github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/repository"
	jadwalPenyuluhanService "github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/service"
	jadwalPosyanduController "github.com/fathoor/posyandu-api/module/jadwal-posyandu/controller"
	jadwalPosyanduRepository "github.com/fathoor/posyandu-api/module/jadwal-posyandu/repository"
	jadwalPosyanduService "github.com/fathoor/posyandu-api/module/jadwal-posyandu/service"
	pemeriksaanController "github.com/fathoor/posyandu-api/module/pemeriksaan/controller"
	pemeriksaanRepository "github.com/fathoor/posyandu-api/module/pemeriksaan/repository"
	pemeriksaanService "github.com/fathoor/posyandu-api/module/pemeriksaan/service"
	pengampuController "github.com/fathoor/posyandu-api/module/pengampu/controller"
	pengampuRepository "github.com/fathoor/posyandu-api/module/pengampu/repository"
	pengampuService "github.com/fathoor/posyandu-api/module/pengampu/service"
	posyanduController "github.com/fathoor/posyandu-api/module/posyandu/controller"
	posyanduRepository "github.com/fathoor/posyandu-api/module/posyandu/repository"
	posyanduService "github.com/fathoor/posyandu-api/module/posyandu/service"
	remajaController "github.com/fathoor/posyandu-api/module/remaja/controller"
	remajaRepository "github.com/fathoor/posyandu-api/module/remaja/repository"
	remajaService "github.com/fathoor/posyandu-api/module/remaja/service"
	controller2 "github.com/fathoor/posyandu-api/module/threshold/controller"
	thresholdRepository "github.com/fathoor/posyandu-api/module/threshold/repository"
	service2 "github.com/fathoor/posyandu-api/module/threshold/service"
	userController "github.com/fathoor/posyandu-api/module/user/controller"
	userRepository "github.com/fathoor/posyandu-api/module/user/repository"
	userService "github.com/fathoor/posyandu-api/module/user/service"
	"github.com/gofiber/fiber/v2"
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
	ProvideThreshold(app, db)
	ProvideFile(app)
	ProvideHome(app, db)
	ProvideChat(app, db)
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
	pemeriksaanRepo := pemeriksaanRepository.ProvidePemeriksaanRepository(db)
	service := remajaService.ProvideRemajaService(&remajaRepo, &posyanduRepo, &userRepo, &pemeriksaanRepo)
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

func ProvideThreshold(app *fiber.App, db *gorm.DB) {
	repository := thresholdRepository.ProvideThresholdRepository(db)
	service := service2.ProvideThresholdService(&repository)
	controller := controller2.ProvideThresholdController(&service)

	controller.Route(app)
}

func ProvideFile(app *fiber.App) {
	service := fileService.ProvideFileService()
	controller := fileController.ProvideFileController(&service)

	controller.Route(app)
}

func ProvideHome(app *fiber.App, db *gorm.DB) {
	userRepo := userRepository.ProvideUserRepository(db)
	bidanRepo := bidanRepository.ProvideBidanRepository(db)
	remajaRepo := remajaRepository.ProvideRemajaRepository(db)
	pengampuRepo := pengampuRepository.ProvidePengampuRepository(db)
	posyanduRepo := posyanduRepository.ProvidePosyanduRepository(db)
	pemeriksaanRepo := pemeriksaanRepository.ProvidePemeriksaanRepository(db)
	jadwalPosyanduRepo := jadwalPosyanduRepository.ProvideJadwalPosyanduRepository(db)
	jadwalPenyuluhanRepo := jadwalPenyuluhanRepository.ProvideJadwalPenyuluhanRepository(db)
	service := homeService.ProvideHomeService(&userRepo, &bidanRepo, &remajaRepo, &pengampuRepo, &posyanduRepo, &pemeriksaanRepo, &jadwalPosyanduRepo, &jadwalPenyuluhanRepo)
	controller := homeController.ProvideHomeController(&service)

	controller.Route(app)
}

func ProvideChat(app *fiber.App, db *gorm.DB) {
	chatRepo := chatRepository.ProvideChatRepository(db)
	userRepo := userRepository.ProvideUserRepository(db)
	service := chatService.ProvideChatService(&chatRepo, &userRepo)
	controller := chatController.ProvideChatController(&service)

	controller.Route(app)
}
