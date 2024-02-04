package validation

import (
	"github.com/fathoor/posyandu-api/module/user/model"
	"github.com/go-playground/validator/v10"
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

func ValidateUserUpdateAuthRequest(request *model.UserUpdateAuthRequest) error {
	return validator.New().Struct(request)
}
