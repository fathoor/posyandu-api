package service

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/core/helper"
	"github.com/itsLeonB/posyandu-api/module/user/entity"
	"github.com/itsLeonB/posyandu-api/module/user/model"
	"github.com/itsLeonB/posyandu-api/module/user/repository"
	"github.com/itsLeonB/posyandu-api/module/user/validation"
	"time"
)

type userServiceImpl struct {
	repository.UserRepository
}

func (service *userServiceImpl) Login(request *model.LoginRequest) (model.LoginResponse, error) {
	valid := validation.ValidateLoginRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	user, err := service.UserRepository.FindByUsername(request.Username)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Account not found",
		})
	}

	if !helper.DecryptPassword(user.Password, request.Password) {
		panic(exception.UnauthorizedError{
			Message: "Wrong password",
		})
	}

	token, err := helper.GenerateJWT(user.Username, user.Role)
	exception.PanicIfNeeded(err)

	response := model.LoginResponse{
		Token:     token,
		Type:      "Bearer",
		ExpiresAt: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (service *userServiceImpl) Register(request *model.UserRegisterRequest) (model.UserResponse, error) {
	valid := validation.ValidateUserRegisterRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	encrypted, err := helper.EncryptPassword(request.Password)
	exception.PanicIfNeeded(err)

	user := entity.User{
		Nama:         request.Nama,
		Email:        request.Email,
		Username:     request.Username,
		Password:     string(encrypted),
		NIK:          request.NIK,
		TempatLahir:  request.TempatLahir,
		TanggalLahir: request.TanggalLahir,
		Alamat:       request.Alamat,
		Provinsi:     request.Provinsi,
		Kota:         request.Kota,
		Kecamatan:    request.Kecamatan,
		Kelurahan:    request.Kelurahan,
		KodePos:      request.KodePos,
		RT:           request.RT,
		RW:           request.RW,
		Telepon:      request.Telepon,
		Role:         request.Role,
	}

	err = service.UserRepository.Insert(&user)
	exception.PanicIfNeeded(err)

	response := model.UserResponse{
		Nama:         user.Nama,
		Email:        user.Email,
		Username:     user.Username,
		NIK:          user.NIK,
		TempatLahir:  user.TempatLahir,
		TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
		Alamat:       user.Alamat,
		Provinsi:     user.Provinsi,
		Kota:         user.Kota,
		Kecamatan:    user.Kecamatan,
		Kelurahan:    user.Kelurahan,
		KodePos:      user.KodePos,
		RT:           user.RT,
		RW:           user.RW,
		Telepon:      user.Telepon,
		Role:         user.Role,
	}

	return response, nil
}

func (service *userServiceImpl) GetAll() ([]model.UserResponse, error) {
	user, err := service.UserRepository.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.UserResponse, len(user))
	for i, user := range user {
		response[i] = model.UserResponse{
			Nama:         user.Nama,
			Email:        user.Email,
			Username:     user.Username,
			NIK:          user.NIK,
			TempatLahir:  user.TempatLahir,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Alamat:       user.Alamat,
			Provinsi:     user.Provinsi,
			Kota:         user.Kota,
			Kecamatan:    user.Kecamatan,
			Kelurahan:    user.Kelurahan,
			KodePos:      user.KodePos,
			RT:           user.RT,
			RW:           user.RW,
			Telepon:      user.Telepon,
			Role:         user.Role,
		}
	}

	return response, err
}

func (service *userServiceImpl) GetByRole(role string) ([]model.UserResponse, error) {
	user, err := service.UserRepository.FindByRole(role)
	exception.PanicIfNeeded(err)

	response := make([]model.UserResponse, len(user))
	for i, user := range user {
		response[i] = model.UserResponse{
			Nama:         user.Nama,
			Email:        user.Email,
			Username:     user.Username,
			NIK:          user.NIK,
			TempatLahir:  user.TempatLahir,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Alamat:       user.Alamat,
			Provinsi:     user.Provinsi,
			Kota:         user.Kota,
			Kecamatan:    user.Kecamatan,
			Kelurahan:    user.Kelurahan,
			KodePos:      user.KodePos,
			RT:           user.RT,
			RW:           user.RW,
			Telepon:      user.Telepon,
			Role:         user.Role,
		}
	}

	return response, err
}

func (service *userServiceImpl) GetByID(id int) (model.UserResponse, error) {
	user, err := service.UserRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	response := model.UserResponse{
		Nama:         user.Nama,
		Email:        user.Email,
		Username:     user.Username,
		NIK:          user.NIK,
		TempatLahir:  user.TempatLahir,
		TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
		Alamat:       user.Alamat,
		Provinsi:     user.Provinsi,
		Kota:         user.Kota,
		Kecamatan:    user.Kecamatan,
		Kelurahan:    user.Kelurahan,
		KodePos:      user.KodePos,
		RT:           user.RT,
		RW:           user.RW,
		Telepon:      user.Telepon,
		Role:         user.Role,
	}

	return response, err
}

func (service *userServiceImpl) Update(id int, request *model.UserUpdateRequest) (model.UserResponse, error) {
	valid := validation.ValidateUserUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	user, err := service.UserRepository.FindByID(id)
	exception.PanicIfNeeded(err)

	if user != (entity.User{}) {
		user.Nama = request.Nama
		user.Email = request.Email
		user.Username = request.Username
		user.Alamat = request.Alamat
		user.Provinsi = request.Provinsi
		user.Kota = request.Kota
		user.Kecamatan = request.Kecamatan
		user.Kelurahan = request.Kelurahan
		user.KodePos = request.KodePos
		user.RT = request.RT
		user.RW = request.RW
		user.Telepon = request.Telepon
	}

	err = service.UserRepository.Save(&user)

	response := model.UserResponse{
		Nama:         user.Nama,
		Email:        user.Email,
		Username:     user.Username,
		NIK:          user.NIK,
		TempatLahir:  user.TempatLahir,
		TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
		Alamat:       user.Alamat,
		Provinsi:     user.Provinsi,
		Kota:         user.Kota,
		Kecamatan:    user.Kecamatan,
		Kelurahan:    user.Kelurahan,
		KodePos:      user.KodePos,
		RT:           user.RT,
		RW:           user.RW,
		Telepon:      user.Telepon,
		Role:         user.Role,
	}

	return response, err
}

func (service *userServiceImpl) Delete(id int) error {
	user, err := service.UserRepository.FindByID(id)
	exception.PanicIfNeeded(err)

	return service.UserRepository.Delete(&user)
}

func ProvideUserService(repository *repository.UserRepository) UserService {
	return &userServiceImpl{*repository}
}
