package service

import (
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/module/posyandu/entity"
	"github.com/itsLeonB/posyandu-api/module/posyandu/model"
	"github.com/itsLeonB/posyandu-api/module/posyandu/repository"
	"github.com/itsLeonB/posyandu-api/module/posyandu/validation"
)

type posyanduServiceImpl struct {
	repository.PosyanduRepository
}

func (service *posyanduServiceImpl) Create(request *model.PosyanduCreateRequest) (model.PosyanduResponse, error) {
	valid := validation.ValidatePosyanduCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	posyandu := entity.Posyandu{
		Nama:      request.Nama,
		Alamat:    request.Alamat,
		Provinsi:  request.Provinsi,
		Kota:      request.Kota,
		Kecamatan: request.Kecamatan,
		Kelurahan: request.Kelurahan,
		KodePos:   request.KodePos,
		RT:        request.RT,
		RW:        request.RW,
		Foto:      request.Foto,
	}

	err := service.PosyanduRepository.Insert(&posyandu)
	exception.PanicIfNeeded(err)

	response := model.PosyanduResponse{
		ID:        posyandu.ID,
		Nama:      posyandu.Nama,
		Alamat:    posyandu.Alamat,
		Provinsi:  posyandu.Provinsi,
		Kota:      posyandu.Kota,
		Kecamatan: posyandu.Kecamatan,
		Kelurahan: posyandu.Kelurahan,
		KodePos:   posyandu.KodePos,
		RT:        posyandu.RT,
		RW:        posyandu.RW,
		Foto:      posyandu.Foto,
	}

	return response, nil
}

func (service *posyanduServiceImpl) GetAll() ([]model.PosyanduResponse, error) {
	posyandu, err := service.PosyanduRepository.FindAll()
	exception.PanicIfNeeded(err)

	response := make([]model.PosyanduResponse, len(posyandu))
	for i, posyandu := range posyandu {
		response[i] = model.PosyanduResponse{
			ID:        posyandu.ID,
			Nama:      posyandu.Nama,
			Alamat:    posyandu.Alamat,
			Provinsi:  posyandu.Provinsi,
			Kota:      posyandu.Kota,
			Kecamatan: posyandu.Kecamatan,
			Kelurahan: posyandu.Kelurahan,
			KodePos:   posyandu.KodePos,
			RT:        posyandu.RT,
			RW:        posyandu.RW,
			Foto:      posyandu.Foto,
		}
	}

	return response, nil
}

func (service *posyanduServiceImpl) GetByID(id int) (model.PosyanduResponse, error) {
	posyandu, err := service.PosyanduRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	response := model.PosyanduResponse{
		ID:        posyandu.ID,
		Nama:      posyandu.Nama,
		Alamat:    posyandu.Alamat,
		Provinsi:  posyandu.Provinsi,
		Kota:      posyandu.Kota,
		Kecamatan: posyandu.Kecamatan,
		Kelurahan: posyandu.Kelurahan,
		KodePos:   posyandu.KodePos,
		RT:        posyandu.RT,
		RW:        posyandu.RW,
		Foto:      posyandu.Foto,
	}

	return response, nil
}

func (service *posyanduServiceImpl) Update(id int, request *model.PosyanduUpdateRequest) (model.PosyanduResponse, error) {
	valid := validation.ValidatePosyanduUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	posyandu, err := service.PosyanduRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	if posyandu != (entity.Posyandu{}) {
		posyandu.Nama = request.Nama
		posyandu.Alamat = request.Alamat
		posyandu.Provinsi = request.Provinsi
		posyandu.Kota = request.Kota
		posyandu.Kecamatan = request.Kecamatan
		posyandu.Kelurahan = request.Kelurahan
		posyandu.KodePos = request.KodePos
		posyandu.RT = request.RT
		posyandu.RW = request.RW
		posyandu.Foto = request.Foto
	}

	err = service.PosyanduRepository.Save(&posyandu)
	exception.PanicIfNeeded(err)

	response := model.PosyanduResponse{
		ID:        posyandu.ID,
		Nama:      posyandu.Nama,
		Alamat:    posyandu.Alamat,
		Provinsi:  posyandu.Provinsi,
		Kota:      posyandu.Kota,
		Kecamatan: posyandu.Kecamatan,
		Kelurahan: posyandu.Kelurahan,
		KodePos:   posyandu.KodePos,
		RT:        posyandu.RT,
		RW:        posyandu.RW,
		Foto:      posyandu.Foto,
	}

	return response, nil
}

func (service *posyanduServiceImpl) Delete(id int) error {
	posyandu, err := service.PosyanduRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Posyandu not found",
		})
	}

	return service.PosyanduRepository.Delete(&posyandu)
}

func ProvidePosyanduService(repository *repository.PosyanduRepository) PosyanduService {
	return &posyanduServiceImpl{*repository}
}
