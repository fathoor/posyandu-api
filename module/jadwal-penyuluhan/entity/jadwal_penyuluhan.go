package entity

import (
	posyanduEntity "github.com/itsLeonB/posyandu-api/module/posyandu/entity"
	"time"
)

type JadwalPenyuluhan struct {
	ID           int                     `gorm:"column:id;primaryKey;autoIncrement"`
	PosyanduID   int                     `gorm:"column:posyandu_id;not null"`
	Posyandu     posyanduEntity.Posyandu `gorm:"foreignKey:posyandu_id;references:id"`
	WaktuMulai   time.Time               `gorm:"column:waktu_mulai;not null"`
	WaktuSelesai time.Time               `gorm:"column:waktu_selesai;not null"`
	Title        string                  `gorm:"column:title;not null"`
	Materi       string                  `gorm:"column:materi"`
	Feedback     string                  `gorm:"column:feedback"`
	CreatedAt    time.Time               `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time               `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (JadwalPenyuluhan) TableName() string {
	return "jadwal_penyuluhan"
}
