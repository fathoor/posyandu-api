package service

import (
	"github.com/fathoor/posyandu-api/core/exception"
	"github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/entity"
	"github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/model"
	jadwalPenyuluhanRepository "github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/repository"
	"github.com/fathoor/posyandu-api/module/jadwal-penyuluhan/validation"
	posyanduRepository "github.com/fathoor/posyandu-api/module/posyandu/repository"
	"time"
)

type jadwalPenyuluhanServiceImpl struct {
	jadwalPenyuluhanRepo jadwalPenyuluhanRepository.JadwalPenyuluhanRepository
	posyanduRepo         posyanduRepository.PosyanduRepository
}

func (service *jadwalPenyuluhanServiceImpl) Create(request *model.JadwalPenyuluhanCreateRequest) (model.JadwalPenyuluhanResponse, error) {
	valid := validation.ValidateJadwalPenyuluhanCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jadwalPenyuluhan := entity.JadwalPenyuluhan{
		PosyanduID:   request.PosyanduID,
		WaktuMulai:   request.WaktuMulai,
		WaktuSelesai: request.WaktuSelesai,
		Title:        request.Title,
		Materi:       request.Materi,
		Feedback:     request.Feedback,
	}

	posyandu, err := service.posyanduRepo.FindByID(jadwalPenyuluhan.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	err = service.jadwalPenyuluhanRepo.Insert(&jadwalPenyuluhan)
	exception.PanicIfNeeded(err)

	response := model.JadwalPenyuluhanResponse{
		ID: jadwalPenyuluhan.ID,
		Posyandu: model.JadwalPenyuluhanPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		WaktuMulai:   jadwalPenyuluhan.WaktuMulai.Format("2006-01-02 15:04:05"),
		WaktuSelesai: jadwalPenyuluhan.WaktuSelesai.Format("2006-01-02 15:04:05"),
		Title:        jadwalPenyuluhan.Title,
		Materi:       jadwalPenyuluhan.Materi,
		Feedback:     jadwalPenyuluhan.Feedback,
	}

	return response, nil
}

func (service *jadwalPenyuluhanServiceImpl) GetAll() ([]model.JadwalPenyuluhanResponse, error) {
	jadwalPenyuluhan, err := service.jadwalPenyuluhanRepo.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.JadwalPenyuluhanResponse, len(jadwalPenyuluhan))
	for i, jadwalPenyuluhan := range jadwalPenyuluhan {
		posyandu, err := service.posyanduRepo.FindByID(jadwalPenyuluhan.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		response[i] = model.JadwalPenyuluhanResponse{
			ID: jadwalPenyuluhan.ID,
			Posyandu: model.JadwalPenyuluhanPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			WaktuMulai:   jadwalPenyuluhan.WaktuMulai.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
			WaktuSelesai: jadwalPenyuluhan.WaktuSelesai.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
			Title:        jadwalPenyuluhan.Title,
			Materi:       jadwalPenyuluhan.Materi,
			Feedback:     jadwalPenyuluhan.Feedback,
		}
	}

	return response, nil
}

func (service *jadwalPenyuluhanServiceImpl) GetByPosyanduID(id int) ([]model.JadwalPenyuluhanResponse, error) {
	jadwalPenyuluhan, err := service.jadwalPenyuluhanRepo.FindByPosyanduID(id)
	exception.PanicIfNeeded(err)

	response := make([]model.JadwalPenyuluhanResponse, len(jadwalPenyuluhan))
	for i, jadwalPenyuluhan := range jadwalPenyuluhan {
		posyandu, err := service.posyanduRepo.FindByID(jadwalPenyuluhan.PosyanduID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Posyandu not found",
			})
		}

		response[i] = model.JadwalPenyuluhanResponse{
			ID: jadwalPenyuluhan.ID,
			Posyandu: model.JadwalPenyuluhanPosyanduResponse{
				ID:     posyandu.ID,
				Nama:   posyandu.Nama,
				Alamat: posyandu.Alamat,
				Foto:   posyandu.Foto,
			},
			WaktuMulai:   jadwalPenyuluhan.WaktuMulai.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
			WaktuSelesai: jadwalPenyuluhan.WaktuSelesai.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
			Title:        jadwalPenyuluhan.Title,
			Materi:       jadwalPenyuluhan.Materi,
			Feedback:     jadwalPenyuluhan.Feedback,
		}
	}

	return response, nil
}

func (service *jadwalPenyuluhanServiceImpl) GetByID(id int) (model.JadwalPenyuluhanResponse, error) {
	jadwalPenyuluhan, err := service.jadwalPenyuluhanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal penyuluhan not found",
		})
	}

	posyandu, err := service.posyanduRepo.FindByID(jadwalPenyuluhan.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	response := model.JadwalPenyuluhanResponse{
		ID: jadwalPenyuluhan.ID,
		Posyandu: model.JadwalPenyuluhanPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		WaktuMulai:   jadwalPenyuluhan.WaktuMulai.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
		WaktuSelesai: jadwalPenyuluhan.WaktuSelesai.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
		Title:        jadwalPenyuluhan.Title,
		Materi:       jadwalPenyuluhan.Materi,
		Feedback:     jadwalPenyuluhan.Feedback,
	}

	return response, nil
}

func (service *jadwalPenyuluhanServiceImpl) Update(id int, request *model.JadwalPenyuluhanUpdateRequest) (model.JadwalPenyuluhanResponse, error) {
	valid := validation.ValidateJadwalPenyuluhanUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jadwalPenyuluhan, err := service.jadwalPenyuluhanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal penyuluhan not found",
		})
	}

	if jadwalPenyuluhan != (entity.JadwalPenyuluhan{}) {
		jadwalPenyuluhan.PosyanduID = request.PosyanduID
		jadwalPenyuluhan.WaktuMulai = request.WaktuMulai
		jadwalPenyuluhan.WaktuSelesai = request.WaktuSelesai
		jadwalPenyuluhan.Title = request.Title
		jadwalPenyuluhan.Materi = request.Materi
		jadwalPenyuluhan.Feedback = request.Feedback
	}

	posyandu, err := service.posyanduRepo.FindByID(jadwalPenyuluhan.PosyanduID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	err = service.jadwalPenyuluhanRepo.Save(&jadwalPenyuluhan)
	exception.PanicIfNeeded(err)

	response := model.JadwalPenyuluhanResponse{
		ID: jadwalPenyuluhan.ID,
		Posyandu: model.JadwalPenyuluhanPosyanduResponse{
			ID:     posyandu.ID,
			Nama:   posyandu.Nama,
			Alamat: posyandu.Alamat,
			Foto:   posyandu.Foto,
		},
		WaktuMulai:   jadwalPenyuluhan.WaktuMulai.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
		WaktuSelesai: jadwalPenyuluhan.WaktuSelesai.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
		Title:        jadwalPenyuluhan.Title,
		Materi:       jadwalPenyuluhan.Materi,
		Feedback:     jadwalPenyuluhan.Feedback,
	}

	return response, nil
}

func (service *jadwalPenyuluhanServiceImpl) Delete(id int) error {
	jadwalPenyuluhan, err := service.jadwalPenyuluhanRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal penyuluhan not found",
		})
	}

	return service.jadwalPenyuluhanRepo.Delete(&jadwalPenyuluhan)
}

func ProvideJadwalPenyuluhanService(
	jadwalPenyuluhanRepo *jadwalPenyuluhanRepository.JadwalPenyuluhanRepository,
	posyanduRepo *posyanduRepository.PosyanduRepository,
) JadwalPenyuluhanService {
	return &jadwalPenyuluhanServiceImpl{*jadwalPenyuluhanRepo, *posyanduRepo}
}
