package repository

import (
	"github.com/itsLeonB/posyandu-api/module/remaja/entity"
	"gorm.io/gorm"
)

type remajaRepositoryImpl struct {
	*gorm.DB
}

func (repository *remajaRepositoryImpl) Insert(remaja *entity.Remaja) error {
	return repository.DB.Create(&remaja).Error
}

func (repository *remajaRepositoryImpl) FindAll() ([]entity.Remaja, error) {
	var remaja []entity.Remaja
	err := repository.DB.Find(&remaja).Error

	return remaja, err
}

func (repository *remajaRepositoryImpl) FindByID(id int) (entity.Remaja, error) {
	var remaja entity.Remaja
	err := repository.DB.Take(&remaja, id).Error

	return remaja, err
}

func (repository *remajaRepositoryImpl) FindByUserID(id int) (entity.Remaja, error) {
	var remaja entity.Remaja
	err := repository.DB.Take(&remaja, "user_id = ?", id).Error

	return remaja, err
}

func (repository *remajaRepositoryImpl) Save(remaja *entity.Remaja) error {
	return repository.DB.Save(&remaja).Error
}

func (repository *remajaRepositoryImpl) Delete(remaja *entity.Remaja) error {
	return repository.DB.Delete(&remaja).Error
}

func ProvideRemajaRepository(db *gorm.DB) RemajaRepository {
	return &remajaRepositoryImpl{db}
}
