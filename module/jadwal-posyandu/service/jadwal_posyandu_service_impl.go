package service

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/entity"
	"github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/model"
	jadwalPosyanduRepository "github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/repository"
	"github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/validation"
	posyanduRepository "github.com/itsLeonB/posyandu-api/module/posyandu/repository"
)

type jadwalPosyanduServiceImpl struct {
	jadwalPosyanduRepo jadwalPosyanduRepository.JadwalPosyanduRepository
	posyanduRepo       posyanduRepository.PosyanduRepository
}

func (service *jadwalPosyanduServiceImpl) Create(request *model.JadwalPosyanduCreateRequest) (model.JadwalPosyanduResponse, error) {
	valid := validation.ValidateJadwalPosyanduCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jadwalPosyandu := entity.JadwalPosyandu{
		PosyanduID:   request.PosyanduID,
		WaktuMulai:   request.WaktuMulai,
		WaktuSelesai: request.WaktuSelesai,
	}

	posyandu, err := service.posyanduRepo.FindByID(jadwalPosyandu.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	err = service.jadwalPosyanduRepo.Insert(&jadwalPosyandu)
	exception.PanicIfNeeded(err)

	response := model.JadwalPosyanduResponse{
		ID: jadwalPosyandu.ID,
		Posyandu: model.JadwalPosyanduPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		WaktuMulai:   jadwalPosyandu.WaktuMulai.Format("2006-01-02 15:04:05"),
		WaktuSelesai: jadwalPosyandu.WaktuSelesai.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (service *jadwalPosyanduServiceImpl) GetAll() ([]model.JadwalPosyanduResponse, error) {
	jadwalPosyandu, err := service.jadwalPosyanduRepo.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.JadwalPosyanduResponse, len(jadwalPosyandu))
	for i, jadwalPosyandu := range jadwalPosyandu {
		posyandu, err := service.posyanduRepo.FindByID(jadwalPosyandu.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		response[i] = model.JadwalPosyanduResponse{
			ID: jadwalPosyandu.ID,
			Posyandu: model.JadwalPosyanduPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			WaktuMulai:   jadwalPosyandu.WaktuMulai.Format("2006-01-02 15:04:05"),
			WaktuSelesai: jadwalPosyandu.WaktuSelesai.Format("2006-01-02 15:04:05"),
		}
	}

	return response, nil
}

func (service *jadwalPosyanduServiceImpl) GetByPosyanduID(posyanduID int) ([]model.JadwalPosyanduResponse, error) {
	jadwalPosyandu, err := service.jadwalPosyanduRepo.FindByPosyanduID(posyanduID)
	exception.PanicIfNeeded(err)

	response := make([]model.JadwalPosyanduResponse, len(jadwalPosyandu))
	for i, jadwalPosyandu := range jadwalPosyandu {
		posyandu, err := service.posyanduRepo.FindByID(jadwalPosyandu.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		response[i] = model.JadwalPosyanduResponse{
			ID: jadwalPosyandu.ID,
			Posyandu: model.JadwalPosyanduPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			WaktuMulai:   jadwalPosyandu.WaktuMulai.Format("2006-01-02 15:04:05"),
			WaktuSelesai: jadwalPosyandu.WaktuSelesai.Format("2006-01-02 15:04:05"),
		}
	}

	return response, nil
}

func (service *jadwalPosyanduServiceImpl) GetByID(id int) (model.JadwalPosyanduResponse, error) {
	jadwalPosyandu, err := service.jadwalPosyanduRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal Posyandu not found",
		})
	}

	posyandu, err := service.posyanduRepo.FindByID(jadwalPosyandu.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	response := model.JadwalPosyanduResponse{
		ID: jadwalPosyandu.ID,
		Posyandu: model.JadwalPosyanduPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		WaktuMulai:   jadwalPosyandu.WaktuMulai.Format("2006-01-02 15:04:05"),
		WaktuSelesai: jadwalPosyandu.WaktuSelesai.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (service *jadwalPosyanduServiceImpl) Update(id int, request *model.JadwalPosyanduUpdateRequest) (model.JadwalPosyanduResponse, error) {
	valid := validation.ValidateJadwalPosyanduUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jadwalPosyandu, err := service.jadwalPosyanduRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal Posyandu not found",
		})
	}

	posyandu, err := service.posyanduRepo.FindByID(jadwalPosyandu.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	if jadwalPosyandu != (entity.JadwalPosyandu{}) {
		jadwalPosyandu.PosyanduID = request.PosyanduID
		jadwalPosyandu.WaktuMulai = request.WaktuMulai
		jadwalPosyandu.WaktuSelesai = request.WaktuSelesai
	}

	err = service.jadwalPosyanduRepo.Save(&jadwalPosyandu)
	exception.PanicIfNeeded(err)

	response := model.JadwalPosyanduResponse{
		ID: jadwalPosyandu.ID,
		Posyandu: model.JadwalPosyanduPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		WaktuMulai:   jadwalPosyandu.WaktuMulai.Format("2006-01-02 15:04:05"),
		WaktuSelesai: jadwalPosyandu.WaktuSelesai.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (service *jadwalPosyanduServiceImpl) Delete(id int) error {
	jadwalPosyandu, err := service.jadwalPosyanduRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal Posyandu not found",
		})
	}

	return service.jadwalPosyanduRepo.Delete(&jadwalPosyandu)
}

func ProvideJadwalPosyanduService(
	jadwalPosyanduRepo *jadwalPosyanduRepository.JadwalPosyanduRepository,
	posyanduRepo *posyanduRepository.PosyanduRepository,
) JadwalPosyanduService {
	return &jadwalPosyanduServiceImpl{*jadwalPosyanduRepo, *posyanduRepo}
}
