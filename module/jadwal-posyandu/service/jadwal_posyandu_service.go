package service

import "github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/model"

type JadwalPosyanduService interface {
	Create(request *model.JadwalPosyanduCreateRequest) (model.JadwalPosyanduResponse, error)
	GetAll() ([]model.JadwalPosyanduResponse, error)
	GetByID(id int) (model.JadwalPosyanduResponse, error)
	Update(id int, request *model.JadwalPosyanduUpdateRequest) (model.JadwalPosyanduResponse, error)
	Delete(id int) error
}
