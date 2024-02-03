package service

import "github.com/itsLeonB/posyandu-api/module/threshold/model"

type ThresholdService interface {
	Create(request *model.ThresholdCreateRequest) (model.ThresholdResponse, error)
	GetAll() ([]model.ThresholdResponse, error)
	GetByParameter(parameter string) (model.ThresholdResponse, error)
	Update(parameter string, request *model.ThresholdUpdateRequest) (model.ThresholdResponse, error)
	Delete(parameter string) error
}
