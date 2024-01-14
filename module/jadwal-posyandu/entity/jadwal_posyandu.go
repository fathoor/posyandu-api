package entity

import (
	posyanduEntity "github.com/itsLeonB/posyandu-api/module/posyandu/entity"
	"time"
)

type JadwalPosyandu struct {
	ID           int                     `gorm:"column:id;primaryKey;autoIncrement"`
	PosyanduID   int                     `gorm:"column:posyandu_id;not null"`
	Posyandu     posyanduEntity.Posyandu `gorm:"foreignKey:posyandu_id;references:id"`
	WaktuMulai   time.Time               `gorm:"column:waktu_mulai;not null"`
	WaktuSelesai time.Time               `gorm:"column:waktu_selesai;not null"`
	CreatedAt    time.Time               `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time               `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (JadwalPosyandu) TableName() string {
	return "jadwal_posyandu"
}
