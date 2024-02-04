package repository

import (
	"github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/entity"
	"gorm.io/gorm"
)

type jadwalPenyuluhanRepositoryImpl struct {
	*gorm.DB
}

func (repository *jadwalPenyuluhanRepositoryImpl) Insert(jadwalPenyuluhan *entity.JadwalPenyuluhan) error {
	return repository.DB.Create(&jadwalPenyuluhan).Error
}

func (repository *jadwalPenyuluhanRepositoryImpl) FindAll() ([]entity.JadwalPenyuluhan, error) {
	var jadwalPenyuluhan []entity.JadwalPenyuluhan
	err := repository.DB.Find(&jadwalPenyuluhan).Error

	return jadwalPenyuluhan, err
}

func (repository *jadwalPenyuluhanRepositoryImpl) FindByPosyanduID(id int) ([]entity.JadwalPenyuluhan, error) {
	var jadwalPenyuluhan []entity.JadwalPenyuluhan
	err := repository.DB.Find(&jadwalPenyuluhan, "posyandu_id = ?", id).Error

	return jadwalPenyuluhan, err
}

func (repository *jadwalPenyuluhanRepositoryImpl) FindByID(id int) (entity.JadwalPenyuluhan, error) {
	var jadwalPenyuluhan entity.JadwalPenyuluhan
	err := repository.DB.Take(&jadwalPenyuluhan, id).Error

	return jadwalPenyuluhan, err
}

func (repository *jadwalPenyuluhanRepositoryImpl) Save(jadwalPenyuluhan *entity.JadwalPenyuluhan) error {
	return repository.DB.Save(&jadwalPenyuluhan).Error
}

func (repository *jadwalPenyuluhanRepositoryImpl) Delete(jadwalPenyuluhan *entity.JadwalPenyuluhan) error {
	return repository.DB.Delete(&jadwalPenyuluhan).Error
}

func ProvideJadwalPenyuluhanRepository(db *gorm.DB) JadwalPenyuluhanRepository {
	return &jadwalPenyuluhanRepositoryImpl{db}
}
