package repository

import "github.com/itsLeonB/posyandu-api/module/jadwal-posyandu/entity"

type JadwalPosyanduRepository interface {
	Insert(jadwalPosyandu *entity.JadwalPosyandu) error
	FindAll() ([]entity.JadwalPosyandu, error)
	FindByPosyanduID(id int) ([]entity.JadwalPosyandu, error)
	FindByID(id int) (entity.JadwalPosyandu, error)
	Save(jadwalPosyandu *entity.JadwalPosyandu) error
	Delete(jadwalPosyandu *entity.JadwalPosyandu) error
}
