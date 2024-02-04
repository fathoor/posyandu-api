package service

import "github.com/fathoor/posyandu-api/module/pengampu/model"

type PengampuService interface {
	Create(request *model.PengampuCreateRequest) (model.PengampuResponse, error)
	GetAll() ([]model.PengampuResponse, error)
	GetByID(id int) ([]model.PengampuResponse, error)
	Update(request *model.PengampuUpdateRequest) (model.PengampuResponse, error)
	Delete(id, pid int) error
}
