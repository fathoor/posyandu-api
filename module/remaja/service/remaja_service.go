package service

import "github.com/fathoor/posyandu-api/module/remaja/model"

type RemajaService interface {
	Create(request *model.RemajaCreateRequest) (model.RemajaResponse, error)
	GetAll() ([]model.RemajaResponse, error)
	GetAllKader() ([]model.RemajaResponse, error)
	GetByPosyanduID(id int) ([]model.RemajaPemeriksaanResponse, error)
	GetByID(id int) (model.RemajaResponse, error)
	Update(id int, request *model.RemajaUpdateRequest) (model.RemajaResponse, error)
	UpdateKader(id int, request *model.RemajaUpdateKaderRequest) (model.RemajaResponse, error)
	Delete(id int) error
}
