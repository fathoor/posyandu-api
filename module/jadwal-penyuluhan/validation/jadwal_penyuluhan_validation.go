package validation

import (
	"github.com/fathoor/posyandu-api/core/validation"
	"github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/model"
)

func ValidateJadwalPenyuluhanCreateRequest(request *model.JadwalPenyuluhanCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateJadwalPenyuluhanUpdateRequest(request *model.JadwalPenyuluhanUpdateRequest) error {
	return validation.Validator.Struct(request)
}
