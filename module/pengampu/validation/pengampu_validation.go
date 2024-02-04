package validation

import (
	"github.com/fathoor/posyandu-api/core/validation"
	"github.com/fathoor/posyandu-api/module/pengampu/model"
)

func ValidatePengampuCreateRequest(request *model.PengampuCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidatePengampuUpdateRequest(request *model.PengampuUpdateRequest) error {
	return validation.Validator.Struct(request)
}
