package repository

import "github.com/itsLeonB/posyandu-api/module/pengampu/entity"

type PengampuRepository interface {
	Insert(pengampu *entity.Pengampu) error
	FindAll() ([]entity.Pengampu, error)
	FindByID(id int) (entity.Pengampu, error)
	FindByBidanID(id int) (entity.Pengampu, error)
	Save(pengampu *entity.Pengampu) error
	Delete(pengampu *entity.Pengampu) error
}
