package repository

import "github.com/fathoor/posyandu-api/module/posyandu/entity"

type PosyanduRepository interface {
	Insert(posyandu *entity.Posyandu) error
	FindAll() ([]entity.Posyandu, error)
	FindByID(id int) (entity.Posyandu, error)
	Save(posyandu *entity.Posyandu) error
	Delete(posyandu *entity.Posyandu) error
}
