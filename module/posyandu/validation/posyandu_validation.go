package validation

import (
	"github.com/itsLeonB/posyandu-api/core/validation"
	"github.com/itsLeonB/posyandu-api/module/posyandu/model"
)

func ValidatePosyanduCreateRequest(request *model.PosyanduCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidatePosyanduUpdateRequest(request *model.PosyanduUpdateRequest) error {
	return validation.Validator.Struct(request)
}
