package repository

import "github.com/fathoor/posyandu-api/module/pemeriksaan/entity"

type PemeriksaanRepository interface {
	Insert(pemeriksaan *entity.Pemeriksaan) error
	FindAll() ([]entity.Pemeriksaan, error)
	FindAllByRemajaID(id int) ([]entity.Pemeriksaan, error)
	FindAllByPosyanduID(id int) ([]entity.Pemeriksaan, error)
	FindByID(id int) (entity.Pemeriksaan, error)
	FindLastByRemajaID(id int) (entity.Pemeriksaan, error)
	Save(pemeriksaan *entity.Pemeriksaan) error
	Delete(pemeriksaan *entity.Pemeriksaan) error
}
