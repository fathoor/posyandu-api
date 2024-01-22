package service

import "github.com/itsLeonB/posyandu-api/module/jadwal-penyuluhan/model"

type JadwalPenyuluhanService interface {
	Create(request *model.JadwalPenyuluhanCreateRequest) (model.JadwalPenyuluhanResponse, error)
	GetAll() ([]model.JadwalPenyuluhanResponse, error)
	GetByPosyanduID(id int) ([]model.JadwalPenyuluhanResponse, error)
	GetByID(id int) (model.JadwalPenyuluhanResponse, error)
	Update(id int, request *model.JadwalPenyuluhanUpdateRequest) (model.JadwalPenyuluhanResponse, error)
	Delete(id int) error
}
