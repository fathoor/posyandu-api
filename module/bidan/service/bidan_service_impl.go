package service

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/module/bidan/entity"
	"github.com/itsLeonB/posyandu-api/module/bidan/model"
	bidanRepository "github.com/itsLeonB/posyandu-api/module/bidan/repository"
	"github.com/itsLeonB/posyandu-api/module/bidan/validation"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
)

type bidanServiceImpl struct {
	bidanRepo bidanRepository.BidanRepository
	userRepo  userRepository.UserRepository
}

func (service *bidanServiceImpl) Create(request *model.BidanCreateRequest) (model.BidanResponse, error) {
	valid := validation.ValidateBidanCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	bidan := entity.Bidan{
		UserID:  request.UserID,
		Jabatan: request.Jabatan,
	}

	user, err := service.userRepo.FindByID(bidan.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	err = service.bidanRepo.Insert(&bidan)
	exception.PanicIfNeeded(err)

	response := model.BidanResponse{
		ID: bidan.ID,
		User: model.BidanUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
			Role:         user.Role,
		},
		Jabatan: bidan.Jabatan,
	}

	return response, nil
}

func (service *bidanServiceImpl) GetAll() ([]model.BidanResponse, error) {
	bidan, err := service.bidanRepo.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.BidanResponse, len(bidan))
	for i, bidan := range bidan {
		user, err := service.userRepo.FindByID(bidan.UserID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "User not found",
			})
		}

		response[i] = model.BidanResponse{
			ID: bidan.ID,
			User: model.BidanUserResponse{
				ID:           user.ID,
				Nama:         user.Nama,
				NIK:          user.NIK,
				TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
				Foto:         user.Foto,
				Role:         user.Role,
			},
			Jabatan: bidan.Jabatan,
		}
	}

	return response, nil
}

func (service *bidanServiceImpl) GetByID(id int) (model.BidanResponse, error) {
	bidan, err := service.bidanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Bidan not found",
		})
	}

	user, err := service.userRepo.FindByID(bidan.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	response := model.BidanResponse{
		ID: bidan.ID,
		User: model.BidanUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
			Role:         user.Role,
		},
		Jabatan: bidan.Jabatan,
	}

	return response, nil
}

func (service *bidanServiceImpl) Update(id int, request *model.BidanUpdateRequest) (model.BidanResponse, error) {
	valid := validation.ValidateBidanUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	bidan, err := service.bidanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Bidan not found",
		})
	}

	bidan.Jabatan = request.Jabatan

	err = service.bidanRepo.Save(&bidan)
	exception.PanicIfNeeded(err)

	user, err := service.userRepo.FindByID(bidan.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	response := model.BidanResponse{
		ID: bidan.ID,
		User: model.BidanUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
			Role:         user.Role,
		},
		Jabatan: bidan.Jabatan,
	}

	return response, nil
}

func (service *bidanServiceImpl) Delete(id int) error {
	bidan, err := service.bidanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Bidan not found",
		})
	}

	return service.bidanRepo.Delete(&bidan)
}

func ProvideBidanService(bidanRepo *bidanRepository.BidanRepository, userRepo *userRepository.UserRepository) BidanService {
	return &bidanServiceImpl{*bidanRepo, *userRepo}
}
