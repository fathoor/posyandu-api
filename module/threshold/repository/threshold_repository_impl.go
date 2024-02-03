package repository

import (
	"github.com/itsLeonB/posyandu-api/module/threshold/entity"
	"gorm.io/gorm"
)

type thresholdRepositoryImpl struct {
	*gorm.DB
}

func (repository *thresholdRepositoryImpl) Insert(threshold *entity.Threshold) error {
	return repository.DB.Create(&threshold).Error
}

func (repository *thresholdRepositoryImpl) FindAll() ([]entity.Threshold, error) {
	var threshold []entity.Threshold
	err := repository.DB.Find(&threshold).Error

	return threshold, err
}

func (repository *thresholdRepositoryImpl) FindByParameter(parameter string) (entity.Threshold, error) {
	var threshold entity.Threshold
	err := repository.DB.Find(&threshold, "parameter = ?", parameter).Error

	return threshold, err
}

func (repository *thresholdRepositoryImpl) Save(threshold *entity.Threshold) error {
	return repository.DB.Save(&threshold).Error
}

func (repository *thresholdRepositoryImpl) Delete(threshold *entity.Threshold) error {
	return repository.DB.Delete(&threshold).Error
}

func ProvideThresholdRepository(db *gorm.DB) ThresholdRepository {
	return &thresholdRepositoryImpl{db}
}
