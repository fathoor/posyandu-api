package service

import "github.com/itsLeonB/posyandu-api/module/pemeriksaan/model"

type PemeriksaanService interface {
	Create(request *model.PemeriksaanCreateRequest) (model.PemeriksaanResponse, error)
	GetAll() ([]model.PemeriksaanResponse, error)
	GetByID(id int) (model.PemeriksaanResponse, error)
	Update(id int, request *model.PemeriksaanUpdateRequest) (model.PemeriksaanResponse, error)
	Delete(id int) error
}
