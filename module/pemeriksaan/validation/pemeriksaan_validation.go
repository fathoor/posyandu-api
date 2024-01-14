package validation

import (
	"github.com/itsLeonB/posyandu-api/core/validation"
	"github.com/itsLeonB/posyandu-api/module/pemeriksaan/model"
)

func ValidatePemeriksaanCreateRequest(request *model.PemeriksaanCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidatePemeriksaanUpdateRequest(request *model.PemeriksaanUpdateRequest) error {
	return validation.Validator.Struct(request)
}
