package validation

import (
	"github.com/itsLeonB/posyandu-api/core/validation"
	"github.com/itsLeonB/posyandu-api/module/remaja/model"
)

func ValidateRemajaCreateRequest(request *model.RemajaCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateRemajaUpdateRequest(request *model.RemajaUpdateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateRemajaUpdateKaderRequest(request *model.RemajaUpdateKaderRequest) error {
	return validation.Validator.Struct(request)
}
