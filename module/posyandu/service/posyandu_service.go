package service

import "github.com/fathoor/posyandu-api/module/posyandu/model"

type PosyanduService interface {
	Create(request *model.PosyanduCreateRequest) (model.PosyanduResponse, error)
	GetAll() ([]model.PosyanduResponse, error)
	GetByID(id int) (model.PosyanduResponse, error)
	Update(id int, request *model.PosyanduUpdateRequest) (model.PosyanduResponse, error)
	Delete(id int) error
}
