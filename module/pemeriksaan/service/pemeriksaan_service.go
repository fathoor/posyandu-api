package service

import "github.com/fathoor/posyandu-api/module/pemeriksaan/model"

type PemeriksaanService interface {
	Create(request *model.PemeriksaanCreateRequest) (model.PemeriksaanResponse, error)
	CreateKader(request *model.PemeriksaanCreateKaderRequest) (model.PemeriksaanResponse, error)
	GetAll() ([]model.PemeriksaanResponse, error)
	GetByRemajaUserID(id int) ([]model.PemeriksaanResponse, error)
	GetByID(id int) (model.PemeriksaanResponse, error)
	Update(id int, request *model.PemeriksaanUpdateRequest) (model.PemeriksaanResponse, error)
	Delete(id int) error
}
