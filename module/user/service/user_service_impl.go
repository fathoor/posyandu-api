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

	token, err := helper.GenerateJWT(user.ID, user.Role)
	exception.PanicIfNeeded(err)

	response := model.LoginResponse{
		Token:     token,
		Role:      user.Role,
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

	if request.Role == "admin" {
		panic(exception.ForbiddenError{
			Message: "Forbidden to register admin",
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
		Foto:         request.Foto,
		Role:         request.Role,
	}

	err = service.UserRepository.Insert(&user)
	exception.PanicIfNeeded(err)

	response := model.UserResponse{
		ID:           user.ID,
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
		Foto:         user.Foto,
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
			ID:           user.ID,
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
			Foto:         user.Foto,
			Role:         user.Role,
		}
	}

	return response, nil
}

func (service *userServiceImpl) GetByRole(role string) ([]model.UserResponse, error) {
	user, err := service.UserRepository.FindByRole(role)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	response := make([]model.UserResponse, len(user))
	for i, user := range user {
		response[i] = model.UserResponse{
			ID:           user.ID,
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
			Foto:         user.Foto,
			Role:         user.Role,
		}
	}

	return response, nil
}

func (service *userServiceImpl) GetByID(id int) (model.UserResponse, error) {
	user, err := service.UserRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	response := model.UserResponse{
		ID:           user.ID,
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
		Foto:         user.Foto,
		Role:         user.Role,
	}

	return response, nil
}

func (service *userServiceImpl) Update(id int, request *model.UserUpdateRequest) (model.UserResponse, error) {
	valid := validation.ValidateUserUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	user, err := service.UserRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

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
		user.Foto = request.Foto
	}

	err = service.UserRepository.Save(&user)
	exception.PanicIfNeeded(err)

	response := model.UserResponse{
		ID:           user.ID,
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
		Foto:         user.Foto,
		Role:         user.Role,
	}

	return response, nil
}

func (service *userServiceImpl) UpdateAuth(id int, request *model.UserUpdateAuthRequest) error {
	valid := validation.ValidateUserUpdateAuthRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	user, err := service.UserRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	if user != (entity.User{}) {
		if !helper.DecryptPassword(user.Password, request.Password) {
			panic(exception.UnauthorizedError{
				Message: "Wrong password",
			})
		}

		encrypted, err := helper.EncryptPassword(request.NewPassword)
		exception.PanicIfNeeded(err)

		user.Username = request.Username
		user.Password = string(encrypted)
	}

	return service.UserRepository.Save(&user)
}

func (service *userServiceImpl) Delete(id int) error {
	user, err := service.UserRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	return service.UserRepository.Delete(&user)
}

func ProvideUserService(repository *repository.UserRepository) UserService {
	return &userServiceImpl{*repository}
}
