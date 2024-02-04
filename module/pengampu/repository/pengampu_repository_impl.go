package repository

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/module/pengampu/entity"
	"gorm.io/gorm"
)

type pengampuRepositoryImpl struct {
	*gorm.DB
}

func (repository *pengampuRepositoryImpl) Insert(pengampu *entity.Pengampu) error {
	if pengampu.Active && repository.DB.Find(&entity.Pengampu{}, "bidan_id = ?", pengampu.BidanID).RowsAffected > 0 {
		err := repository.DB.Model(&entity.Pengampu{}).Where("bidan_id = ?", pengampu.BidanID).Update("active", false).Error
		exception.PanicIfNeeded(err)
	}

	return repository.DB.Create(&pengampu).Error
}

func (repository *pengampuRepositoryImpl) FindAll() ([]entity.Pengampu, error) {
	var pengampu []entity.Pengampu
	err := repository.DB.Find(&pengampu).Error

	return pengampu, err
}

func (repository *pengampuRepositoryImpl) FindByID(id int) ([]entity.Pengampu, error) {
	var pengampu []entity.Pengampu
	err := repository.DB.Find(&pengampu, "bidan_id = ?", id).Error

	return pengampu, err
}

func (repository *pengampuRepositoryImpl) FindByActivePosyanduID(id int) ([]entity.Pengampu, error) {
	var pengampu []entity.Pengampu
	err := repository.DB.Find(&pengampu, "posyandu_id = ? AND active = ?", id, true).Error

	return pengampu, err
}

func (repository *pengampuRepositoryImpl) FindByBidanAndPosyanduID(bidanID int, posyanduID int) (entity.Pengampu, error) {
	var pengampu entity.Pengampu
	err := repository.DB.Take(&pengampu, "bidan_id = ? AND posyandu_id = ?", bidanID, posyanduID).Error

	return pengampu, err
}

func (repository *pengampuRepositoryImpl) FindByActiveBidanID(id int) (entity.Pengampu, error) {
	var pengampu entity.Pengampu
	err := repository.DB.Take(&pengampu, "bidan_id = ? AND active = ?", id, true).Error

	return pengampu, err
}

func (repository *pengampuRepositoryImpl) Save(pengampu *entity.Pengampu) error {
	if pengampu.Active {
		err := repository.DB.Model(&entity.Pengampu{}).Where("bidan_id = ?", pengampu.BidanID).Update("active", false).Error
		exception.PanicIfNeeded(err)
	}

	return repository.DB.Save(&pengampu).Error
}

func (repository *pengampuRepositoryImpl) Delete(pengampu *entity.Pengampu) error {
	return repository.DB.Delete(&pengampu).Error
}

func ProvidePengampuRepository(db *gorm.DB) PengampuRepository {
	return &pengampuRepositoryImpl{db}
}
