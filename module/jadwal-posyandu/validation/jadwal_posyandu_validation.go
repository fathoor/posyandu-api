package validation

import (
	"github.com/fathoor/posyandu-api/core/validation"
	"github.com/fathoor/posyandu-api/module/jadwal-posyandu/model"
)

func ValidateJadwalPosyanduCreateRequest(request *model.JadwalPosyanduCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateJadwalPosyanduUpdateRequest(request *model.JadwalPosyanduUpdateRequest) error {
	return validation.Validator.Struct(request)
}
