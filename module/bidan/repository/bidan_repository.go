package repository

import "github.com/itsLeonB/posyandu-api/module/bidan/entity"

type BidanRepository interface {
	Insert(bidan *entity.Bidan) error
	FindAll() ([]entity.Bidan, error)
	FindByID(id int) (entity.Bidan, error)
	FindByUserID(userID int) (entity.Bidan, error)
	Save(bidan *entity.Bidan) error
	Delete(bidan *entity.Bidan) error
}
