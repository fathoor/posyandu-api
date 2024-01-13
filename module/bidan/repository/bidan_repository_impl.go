package repository

import (
	"github.com/itsLeonB/posyandu-api/module/bidan/entity"
	"gorm.io/gorm"
)

type bidanRepositoryImpl struct {
	*gorm.DB
}

func (repository *bidanRepositoryImpl) Insert(bidan *entity.Bidan) error {
	return repository.DB.Create(&bidan).Error
}

func (repository *bidanRepositoryImpl) FindAll() ([]entity.Bidan, error) {
	var bidan []entity.Bidan
	err := repository.DB.Find(&bidan).Error

	return bidan, err
}

func (repository *bidanRepositoryImpl) FindByID(id int) (entity.Bidan, error) {
	var bidan entity.Bidan
	err := repository.DB.Take(&bidan, id).Error

	return bidan, err
}

func (repository *bidanRepositoryImpl) Save(bidan *entity.Bidan) error {
	return repository.DB.Save(&bidan).Error
}

func (repository *bidanRepositoryImpl) Delete(bidan *entity.Bidan) error {
	return repository.DB.Delete(&bidan).Error
}

func ProvideBidanRepository(db *gorm.DB) BidanRepository {
	return &bidanRepositoryImpl{db}
}
