package repository

import (
	"github.com/itsLeonB/posyandu-api/module/pemeriksaan/entity"
	"gorm.io/gorm"
)

type pemeriksaanRepositoryImpl struct {
	*gorm.DB
}

func (repository *pemeriksaanRepositoryImpl) Insert(pemeriksaan *entity.Pemeriksaan) error {
	return repository.DB.Create(&pemeriksaan).Error
}

func (repository *pemeriksaanRepositoryImpl) FindAll() ([]entity.Pemeriksaan, error) {
	var pemeriksaan []entity.Pemeriksaan
	err := repository.DB.Find(&pemeriksaan).Error

	return pemeriksaan, err
}

func (repository *pemeriksaanRepositoryImpl) FindAllByRemajaID(id int) ([]entity.Pemeriksaan, error) {
	var pemeriksaan []entity.Pemeriksaan
	err := repository.DB.Find(&pemeriksaan, "remaja_id = ?", id).Error

	return pemeriksaan, err
}

func (repository *pemeriksaanRepositoryImpl) FindByID(id int) (entity.Pemeriksaan, error) {
	var pemeriksaan entity.Pemeriksaan
	err := repository.DB.Take(&pemeriksaan, id).Error

	return pemeriksaan, err
}

func (repository *pemeriksaanRepositoryImpl) Save(pemeriksaan *entity.Pemeriksaan) error {
	return repository.DB.Save(&pemeriksaan).Error
}

func (repository *pemeriksaanRepositoryImpl) Delete(pemeriksaan *entity.Pemeriksaan) error {
	return repository.DB.Delete(&pemeriksaan).Error
}

func ProvidePemeriksaanRepository(db *gorm.DB) PemeriksaanRepository {
	return &pemeriksaanRepositoryImpl{db}
}
