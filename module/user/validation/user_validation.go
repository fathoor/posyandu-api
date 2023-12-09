package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/posyandu-api/module/user/model"
)

func ValidateLoginRequest(request *model.LoginRequest) error {
	return validator.New().Struct(request)
}

func ValidateUserRegisterRequest(request *model.UserRegisterRequest) error {
	return validator.New().Struct(request)
}

func ValidateUserUpdateRequest(request *model.UserUpdateRequest) error {
	return validator.New().Struct(request)
}
