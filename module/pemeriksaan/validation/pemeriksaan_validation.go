package validation

import (
	"github.com/fathoor/posyandu-api/core/validation"
	"github.com/fathoor/posyandu-api/module/pemeriksaan/model"
)

func ValidatePemeriksaanCreateRequest(request *model.PemeriksaanCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidatePemeriksaanCreateKaderRequest(request *model.PemeriksaanCreateKaderRequest) error {
	return validation.Validator.Struct(request)
}

func ValidatePemeriksaanUpdateRequest(request *model.PemeriksaanUpdateRequest) error {
	return validation.Validator.Struct(request)
}
