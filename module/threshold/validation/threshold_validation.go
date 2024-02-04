package validation

import (
	"github.com/fathoor/posyandu-api/core/validation"
	"github.com/fathoor/posyandu-api/module/threshold/model"
)

func ValidateThresholdCreateRequest(request *model.ThresholdCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateThresholdUpdateRequest(request *model.ThresholdUpdateRequest) error {
	return validation.Validator.Struct(request)
}
