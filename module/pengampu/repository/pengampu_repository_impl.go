package repository

import (
	"github.com/itsLeonB/posyandu-api/module/pengampu/entity"
	"gorm.io/gorm"
)

type pengampuRepositoryImpl struct {
	*gorm.DB
}

func (repository *pengampuRepositoryImpl) Insert(pengampu *entity.Pengampu) error {
	return repository.DB.Create(&pengampu).Error
}

func (repository *pengampuRepositoryImpl) FindAll() ([]entity.Pengampu, error) {
	var pengampu []entity.Pengampu
	err := repository.DB.Find(&pengampu).Error

	return pengampu, err
}

func (repository *pengampuRepositoryImpl) FindByID(id int) (entity.Pengampu, error) {
	var pengampu entity.Pengampu
	err := repository.DB.Take(&pengampu, id).Error

	return pengampu, err
}

func (repository *pengampuRepositoryImpl) FindByBidanID(id int) (entity.Pengampu, error) {
	var pengampu entity.Pengampu
	err := repository.DB.Last(&pengampu, "bidan_id = ?", id).Error

	return pengampu, err
}

func (repository *pengampuRepositoryImpl) Save(pengampu *entity.Pengampu) error {
	return repository.DB.Save(&pengampu).Error
}

func (repository *pengampuRepositoryImpl) Delete(pengampu *entity.Pengampu) error {
	return repository.DB.Delete(&pengampu).Error
}

func ProvidePengampuRepository(db *gorm.DB) PengampuRepository {
	return &pengampuRepositoryImpl{db}
}
