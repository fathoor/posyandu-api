package service

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	posyanduRepository "github.com/itsLeonB/posyandu-api/module/posyandu/repository"
	"github.com/itsLeonB/posyandu-api/module/remaja/entity"
	"github.com/itsLeonB/posyandu-api/module/remaja/model"
	remajaRepository "github.com/itsLeonB/posyandu-api/module/remaja/repository"
	"github.com/itsLeonB/posyandu-api/module/remaja/validation"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
)

type remajaServiceImpl struct {
	remajaRepo   remajaRepository.RemajaRepository
	posyanduRepo posyanduRepository.PosyanduRepository
	userRepo     userRepository.UserRepository
}

func (service *remajaServiceImpl) Create(request *model.RemajaCreateRequest) (model.RemajaResponse, error) {
	valid := validation.ValidateRemajaCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	remaja := entity.Remaja{
		PosyanduID: request.PosyanduID,
		UserID:     request.UserID,
		NamaAyah:   request.NamaAyah,
		NamaIbu:    request.NamaIbu,
	}

	posyandu, err := service.posyanduRepo.FindByID(remaja.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	user, err := service.userRepo.FindByID(remaja.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	err = service.remajaRepo.Insert(&remaja)
	exception.PanicIfNeeded(err)

	response := model.RemajaResponse{
		ID: remaja.ID,
		Posyandu: model.RemajaPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		User: model.RemajaUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
			Role:         user.Role,
		},
		NamaAyah: remaja.NamaAyah,
		NamaIbu:  remaja.NamaIbu,
		IsKader:  remaja.IsKader,
	}

	return response, nil
}

func (service *remajaServiceImpl) GetAll() ([]model.RemajaResponse, error) {
	remaja, err := service.remajaRepo.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.RemajaResponse, len(remaja))
	for i, remaja := range remaja {
		posyandu, err := service.posyanduRepo.FindByID(remaja.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		user, err := service.userRepo.FindByID(remaja.UserID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "User not found",
			})
		}

		response[i] = model.RemajaResponse{
			ID: remaja.ID,
			Posyandu: model.RemajaPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			User: model.RemajaUserResponse{
				ID:           user.ID,
				Nama:         user.Nama,
				NIK:          user.NIK,
				TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
				Foto:         user.Foto,
				Role:         user.Role,
			},
			NamaAyah: remaja.NamaAyah,
			NamaIbu:  remaja.NamaIbu,
			IsKader:  remaja.IsKader,
		}
	}

	return response, nil
}

func (service *remajaServiceImpl) GetByID(id int) (model.RemajaResponse, error) {
	remaja, err := service.remajaRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Remaja not found",
		})
	}

	posyandu, err := service.posyanduRepo.FindByID(remaja.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	user, err := service.userRepo.FindByID(remaja.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	response := model.RemajaResponse{
		ID: remaja.ID,
		Posyandu: model.RemajaPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		User: model.RemajaUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
			Role:         user.Role,
		},
		NamaAyah: remaja.NamaAyah,
		NamaIbu:  remaja.NamaIbu,
		IsKader:  remaja.IsKader,
	}

	return response, nil
}

func (service *remajaServiceImpl) Update(id int, request *model.RemajaUpdateRequest) (model.RemajaResponse, error) {
	valid := validation.ValidateRemajaUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	remaja, err := service.remajaRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Remaja not found",
		})
	}

	remaja.PosyanduID = request.PosyanduID
	remaja.NamaAyah = request.NamaAyah
	remaja.NamaIbu = request.NamaIbu

	posyandu, err := service.posyanduRepo.FindByID(remaja.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	user, err := service.userRepo.FindByID(remaja.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	err = service.remajaRepo.Save(&remaja)
	exception.PanicIfNeeded(err)

	response := model.RemajaResponse{
		ID: remaja.ID,
		Posyandu: model.RemajaPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		User: model.RemajaUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
			Role:         user.Role,
		},
		NamaAyah: remaja.NamaAyah,
		NamaIbu:  remaja.NamaIbu,
		IsKader:  remaja.IsKader,
	}

	return response, nil
}

func (service *remajaServiceImpl) UpdateKader(id int, request *model.RemajaUpdateKaderRequest) (model.RemajaResponse, error) {
	valid := validation.ValidateRemajaUpdateKaderRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	remaja, err := service.remajaRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Remaja not found",
		})
	}

	remaja.PosyanduID = request.PosyanduID
	remaja.NamaAyah = request.NamaAyah
	remaja.NamaIbu = request.NamaIbu
	remaja.IsKader = request.IsKader

	posyandu, err := service.posyanduRepo.FindByID(remaja.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	user, err := service.userRepo.FindByID(remaja.UserID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "User not found",
		})
	}

	err = service.remajaRepo.Save(&remaja)
	exception.PanicIfNeeded(err)

	response := model.RemajaResponse{
		ID: remaja.ID,
		Posyandu: model.RemajaPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		User: model.RemajaUserResponse{
			ID:           user.ID,
			Nama:         user.Nama,
			NIK:          user.NIK,
			TanggalLahir: user.TanggalLahir.Format("2006-01-02"),
			Foto:         user.Foto,
			Role:         user.Role,
		},
		NamaAyah: remaja.NamaAyah,
		NamaIbu:  remaja.NamaIbu,
		IsKader:  remaja.IsKader,
	}

	return response, nil
}

func (service *remajaServiceImpl) Delete(id int) error {
	remaja, err := service.remajaRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Remaja not found",
		})
	}

	return service.remajaRepo.Delete(&remaja)
}

func ProvideRemajaService(
	remajaRepo *remajaRepository.RemajaRepository,
	posyanduRepo *posyanduRepository.PosyanduRepository,
	userRepo *userRepository.UserRepository,
) RemajaService {
	return &remajaServiceImpl{*remajaRepo, *posyanduRepo, *userRepo}
}
