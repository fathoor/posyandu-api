package repository

import (
	"github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/entity"
	"gorm.io/gorm"
)

type jadwalPosyanduRepositoryImpl struct {
	*gorm.DB
}

func (repository *jadwalPosyanduRepositoryImpl) Insert(jadwalPosyandu *entity.JadwalPosyandu) error {
	return repository.DB.Create(&jadwalPosyandu).Error
}

func (repository *jadwalPosyanduRepositoryImpl) FindAll() ([]entity.JadwalPosyandu, error) {
	var jadwalPosyandu []entity.JadwalPosyandu
	err := repository.DB.Find(&jadwalPosyandu).Error

	return jadwalPosyandu, err
}

func (repository *jadwalPosyanduRepositoryImpl) FindByPosyanduID(id int) ([]entity.JadwalPosyandu, error) {
	var jadwalPosyandu []entity.JadwalPosyandu
	err := repository.DB.Find(&jadwalPosyandu, "posyandu_id = ?", id).Error

	return jadwalPosyandu, err
}

func (repository *jadwalPosyanduRepositoryImpl) FindByID(id int) (entity.JadwalPosyandu, error) {
	var jadwalPosyandu entity.JadwalPosyandu
	err := repository.DB.Take(&jadwalPosyandu, id).Error

	return jadwalPosyandu, err
}

func (repository *jadwalPosyanduRepositoryImpl) Save(jadwalPosyandu *entity.JadwalPosyandu) error {
	return repository.DB.Save(&jadwalPosyandu).Error
}

func (repository *jadwalPosyanduRepositoryImpl) Delete(jadwalPosyandu *entity.JadwalPosyandu) error {
	return repository.DB.Delete(&jadwalPosyandu).Error
}

func ProvideJadwalPosyanduRepository(db *gorm.DB) JadwalPosyanduRepository {
	return &jadwalPosyanduRepositoryImpl{db}
}
