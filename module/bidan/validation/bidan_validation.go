package validation

import (
	"github.com/fathoor/posyandu-api/core/validation"
	"github.com/fathoor/posyandu-api/module/bidan/model"
)

func ValidateBidanCreateRequest(request *model.BidanCreateRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateBidanUpdateRequest(request *model.BidanUpdateRequest) error {
	return validation.Validator.Struct(request)
}
