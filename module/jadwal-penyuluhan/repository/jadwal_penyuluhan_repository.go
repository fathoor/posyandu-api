package repository

import "github.com/itsLeonB/posyandu-api/module/jadwal-penyuluhan/entity"

type JadwalPenyuluhanRepository interface {
	Insert(jadwalPenyuluhan *entity.JadwalPenyuluhan) error
	FindAll() ([]entity.JadwalPenyuluhan, error)
	FindByID(id int) (entity.JadwalPenyuluhan, error)
	Save(jadwalPenyuluhan *entity.JadwalPenyuluhan) error
	Delete(jadwalPenyuluhan *entity.JadwalPenyuluhan) error
}
