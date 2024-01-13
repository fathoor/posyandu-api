package entity

import "time"

type Posyandu struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Nama      string    `gorm:"column:nama;not null"`
	Alamat    string    `gorm:"column:alamat;not null"`
	Provinsi  string    `gorm:"column:provinsi;not null"`
	Kota      string    `gorm:"column:kota;not null"`
	Kecamatan string    `gorm:"column:kecamatan;not null"`
	Kelurahan string    `gorm:"column:kelurahan;not null"`
	KodePos   int       `gorm:"column:kode_pos;not null"`
	RT        int       `gorm:"column:rt;not null"`
	RW        int       `gorm:"column:rw;not null"`
	Foto      string    `gorm:"column:foto;not null;default:'/storage/image/default.jpg'"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Posyandu) TableName() string {
	return "posyandu"
}
