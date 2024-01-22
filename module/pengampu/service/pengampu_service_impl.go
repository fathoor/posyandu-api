package service

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	bidanRepository "github.com/itsLeonB/posyandu-api/module/bidan/repository"
	"github.com/itsLeonB/posyandu-api/module/pengampu/entity"
	"github.com/itsLeonB/posyandu-api/module/pengampu/model"
	pengampuRepository "github.com/itsLeonB/posyandu-api/module/pengampu/repository"
	"github.com/itsLeonB/posyandu-api/module/pengampu/validation"
	posyanduRepository "github.com/itsLeonB/posyandu-api/module/posyandu/repository"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
)

type pengampuServiceImpl struct {
	bidanRepository.BidanRepository
	pengampuRepository.PengampuRepository
	posyanduRepository.PosyanduRepository
	userRepository.UserRepository
}

func (service *pengampuServiceImpl) Create(request *model.PengampuCreateRequest) (model.PengampuResponse, error) {
	valid := validation.ValidatePengampuCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	pengampu := entity.Pengampu{
		BidanID:    request.BidanID,
		PosyanduID: request.PosyanduID,
	}

	bidan, err := service.BidanRepository.FindByID(pengampu.BidanID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Bidan not found",
		})
	}

	user, err := service.UserRepository.FindByID(bidan.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	posyandu, err := service.PosyanduRepository.FindByID(pengampu.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	err = service.PengampuRepository.Insert(&pengampu)
	exception.PanicIfNeeded(err)

	response := model.PengampuResponse{
		Bidan: model.PengampuBidanResponse{
			ID: bidan.ID,
			User: model.PengampuBidanUserResponse{
				ID:           user.ID,
				Nama:         user.Nama,
				NIK:          user.NIK,
				TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
				Foto:         user.Foto,
				Role:         user.Role,
			},
			Jabatan: bidan.Jabatan,
		},
		Posyandu: model.PengampuPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		Active: pengampu.Active,
	}

	return response, nil
}

func (service *pengampuServiceImpl) GetAll() ([]model.PengampuResponse, error) {
	pengampu, err := service.PengampuRepository.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.PengampuResponse, len(pengampu))
	for i, pengampu := range pengampu {
		bidan, err := service.BidanRepository.FindByID(pengampu.BidanID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Bidan not found",
			})
		}

		user, err := service.UserRepository.FindByID(bidan.UserID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "User not found",
			})
		}

		posyandu, err := service.PosyanduRepository.FindByID(pengampu.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		response[i] = model.PengampuResponse{
			Bidan: model.PengampuBidanResponse{
				ID: bidan.ID,
				User: model.PengampuBidanUserResponse{
					ID:           user.ID,
					Nama:         user.Nama,
					NIK:          user.NIK,
					TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
					Foto:         user.Foto,
					Role:         user.Role,
				},
				Jabatan: bidan.Jabatan,
			},
			Posyandu: model.PengampuPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			Active: pengampu.Active,
		}
	}

	return response, nil
}

func (service *pengampuServiceImpl) GetByID(id int) ([]model.PengampuResponse, error) {
	pengampu, err := service.PengampuRepository.FindByID(id)
	exception.PanicIfNeeded(err)

	response := make([]model.PengampuResponse, len(pengampu))
	for i, pengampu := range pengampu {
		bidan, err := service.BidanRepository.FindByID(pengampu.BidanID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Bidan not found",
			})
		}

		user, err := service.UserRepository.FindByID(bidan.UserID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "User not found",
			})
		}

		posyandu, err := service.PosyanduRepository.FindByID(pengampu.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		response[i] = model.PengampuResponse{
			Bidan: model.PengampuBidanResponse{
				ID: bidan.ID,
				User: model.PengampuBidanUserResponse{
					ID:           user.ID,
					Nama:         user.Nama,
					NIK:          user.NIK,
					TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
					Foto:         user.Foto,
					Role:         user.Role,
				},
				Jabatan: bidan.Jabatan,
			},
			Posyandu: model.PengampuPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			Active: pengampu.Active,
		}
	}

	return response, nil
}

func (service *pengampuServiceImpl) Update(request *model.PengampuUpdateRequest) (model.PengampuResponse, error) {
	valid := validation.ValidatePengampuUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	pengampu, err := service.PengampuRepository.FindByBidanAndPosyanduID(request.BidanID, request.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pengampu not found",
		})
	}

	pengampu.BidanID = request.BidanID
	pengampu.PosyanduID = request.PosyanduID
	pengampu.Active = request.Active

	bidan, err := service.BidanRepository.FindByID(pengampu.BidanID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Bidan not found",
		})
	}

	user, err := service.UserRepository.FindByID(bidan.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	posyandu, err := service.PosyanduRepository.FindByID(pengampu.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	err = service.PengampuRepository.Save(&pengampu)
	exception.PanicIfNeeded(err)

	response := model.PengampuResponse{
		Bidan: model.PengampuBidanResponse{
			ID: bidan.ID,
			User: model.PengampuBidanUserResponse{
				ID:           user.ID,
				Nama:         user.Nama,
				NIK:          user.NIK,
				TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
				Foto:         user.Foto,
				Role:         user.Role,
			},
			Jabatan: bidan.Jabatan,
		},
		Posyandu: model.PengampuPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		Active: pengampu.Active,
	}

	return response, nil
}

func (service *pengampuServiceImpl) Delete(id, pid int) error {
	pengampu, err := service.PengampuRepository.FindByBidanAndPosyanduID(id, pid)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pengampu not found",
		})
	}

	return service.PengampuRepository.Delete(&pengampu)
}

func ProvidePengampuService(
	bidanRepository *bidanRepository.BidanRepository,
	pengampuRepository *pengampuRepository.PengampuRepository,
	posyanduRepository *posyanduRepository.PosyanduRepository,
	userRepository *userRepository.UserRepository,
) PengampuService {
	return &pengampuServiceImpl{*bidanRepository, *pengampuRepository, *posyanduRepository, *userRepository}
}
