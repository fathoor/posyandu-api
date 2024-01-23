package entity

import (
	"github.com/itsLeonB/posyandu-api/module/user/entity"
	"time"
)

type Room struct {
	ID          string      `gorm:"column:id;primaryKey"`
	KonsultanID int         `gorm:"column:konsultan_id;not null"`
	Konsultan   entity.User `gorm:"foreignKey:konsultan_id;references:id"`
	PasienID    int         `gorm:"column:pasien_id;not null"`
	Pasien      entity.User `gorm:"foreignKey:pasien_id;references:id"`
	CreatedAt   time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Room) TableName() string {
	return "room"
}
