package repository

import "github.com/itsLeonB/posyandu-api/module/threshold/entity"

type ThresholdRepository interface {
	Insert(threshold *entity.Threshold) error
	FindAll() ([]entity.Threshold, error)
	FindByParameter(parameter string) (entity.Threshold, error)
	Save(threshold *entity.Threshold) error
	Delete(threshold *entity.Threshold) error
}
