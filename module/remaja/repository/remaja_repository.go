package repository

import "github.com/fathoor/posyandu-api/module/remaja/entity"

type RemajaRepository interface {
	Insert(remaja *entity.Remaja) error
	FindAll() ([]entity.Remaja, error)
	FindAllKader() ([]entity.Remaja, error)
	FindByPosyanduID(id int) ([]entity.Remaja, error)
	FindByID(id int) (entity.Remaja, error)
	FindByUserID(id int) (entity.Remaja, error)
	Save(remaja *entity.Remaja) error
	Delete(remaja *entity.Remaja) error
}
