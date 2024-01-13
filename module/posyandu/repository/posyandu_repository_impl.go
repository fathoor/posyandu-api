package repository

import (
	"github.com/itsLeonB/posyandu-api/module/posyandu/entity"
	"gorm.io/gorm"
)

type posyanduRepositoryImpl struct {
	*gorm.DB
}

func (repository *posyanduRepositoryImpl) Insert(posyandu *entity.Posyandu) error {
	return repository.DB.Create(&posyandu).Error
}

func (repository *posyanduRepositoryImpl) FindAll() ([]entity.Posyandu, error) {
	var posyandu []entity.Posyandu
	err := repository.DB.Find(&posyandu).Error

	return posyandu, err
}

func (repository *posyanduRepositoryImpl) FindByID(id int) (entity.Posyandu, error) {
	var posyandu entity.Posyandu
	err := repository.DB.Take(&posyandu, id).Error

	return posyandu, err
}

func (repository *posyanduRepositoryImpl) Save(posyandu *entity.Posyandu) error {
	return repository.DB.Save(&posyandu).Error
}

func (repository *posyanduRepositoryImpl) Delete(posyandu *entity.Posyandu) error {
	return repository.DB.Delete(&posyandu).Error
}

func ProvidePosyanduRepository(db *gorm.DB) PosyanduRepository {
	return &posyanduRepositoryImpl{db}
}
