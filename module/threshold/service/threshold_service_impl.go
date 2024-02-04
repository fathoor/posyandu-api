package service

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/module/threshold/entity"
	"github.com/fathoor/posyandu-api/module/threshold/model"
	"github.com/fathoor/posyandu-api/module/threshold/repository"
	"github.com/fathoor/posyandu-api/module/threshold/validation"
)

type thresholdServiceImpl struct {
	thresholdRepo repository.ThresholdRepository
}

func (service *thresholdServiceImpl) Create(request *model.ThresholdCreateRequest) (model.ThresholdResponse, error) {
	valid := validation.ValidateThresholdCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	threshold := entity.Threshold{
		Parameter: request.Parameter,
		Threshold: request.Threshold,
	}

	err := service.thresholdRepo.Insert(&threshold)
	exception.PanicIfNeeded(err)

	response := model.ThresholdResponse{
		Parameter: threshold.Parameter,
		Threshold: threshold.Threshold,
	}

	return response, nil
}

func (service *thresholdServiceImpl) GetAll() ([]model.ThresholdResponse, error) {
	threshold, err := service.thresholdRepo.FindAll()
	exception.PanicIfNeeded(err)

	thresholdResponse := make([]model.ThresholdResponse, len(threshold))
	for i, threshold := range threshold {
		thresholdResponse[i] = model.ThresholdResponse{
			Parameter: threshold.Parameter,
			Threshold: threshold.Threshold,
		}
	}

	return thresholdResponse, nil
}

func (service *thresholdServiceImpl) GetByParameter(parameter string) (model.ThresholdResponse, error) {
	threshold, err := service.thresholdRepo.FindByParameter(parameter)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Threshold not found",
		})
	}

	response := model.ThresholdResponse{
		Parameter: threshold.Parameter,
		Threshold: threshold.Threshold,
	}

	return response, nil
}

func (service *thresholdServiceImpl) Update(parameter string, request *model.ThresholdUpdateRequest) (model.ThresholdResponse, error) {
	valid := validation.ValidateThresholdUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	threshold, err := service.thresholdRepo.FindByParameter(parameter)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Threshold not found",
		})
	}

	threshold.Threshold = request.Threshold

	err = service.thresholdRepo.Save(&threshold)
	exception.PanicIfNeeded(err)

	response := model.ThresholdResponse{
		Parameter: threshold.Parameter,
		Threshold: threshold.Threshold,
	}

	return response, nil
}

func (service *thresholdServiceImpl) Delete(parameter string) error {
	threshold, err := service.thresholdRepo.FindByParameter(parameter)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Threshold not found",
		})
	}

	return service.thresholdRepo.Delete(&threshold)
}

func ProvideThresholdService(thresholdRepo *repository.ThresholdRepository) ThresholdService {
	return &thresholdServiceImpl{*thresholdRepo}
}
