package service

import "github.com/fathoor/posyandu-api/module/user/model"

type UserService interface {
	Login(request *model.LoginRequest) (model.LoginResponse, error)
	Register(request *model.UserRegisterRequest) (model.UserResponse, error)
	GetAll() ([]model.UserResponse, error)
	GetByRole(role string) ([]model.UserResponse, error)
	GetByID(id int) (model.UserResponse, error)
	Update(id int, request *model.UserUpdateRequest) (model.UserResponse, error)
	UpdateAuth(id int, request *model.UserUpdateAuthRequest) error
	Delete(id int) error
}
