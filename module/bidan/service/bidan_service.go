package service

import "github.com/fathoor/posyandu-api/module/bidan/model"

type BidanService interface {
	Create(request *model.BidanCreateRequest) (model.BidanResponse, error)
	GetAll() ([]model.BidanResponse, error)
	GetByID(id int) (model.BidanResponse, error)
	Update(id int, request *model.BidanUpdateRequest) (model.BidanResponse, error)
	Delete(id int) error
}
