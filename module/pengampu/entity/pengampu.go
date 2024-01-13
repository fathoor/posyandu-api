package entity

import (
	bidanEntity "github.com/itsLeonB/posyandu-api/module/bidan/entity"
	posyanduEntity "github.com/itsLeonB/posyandu-api/module/posyandu/entity"
	"time"
)

type Pengampu struct {
	ID         int                     `gorm:"column:id;primaryKey;autoIncrement"`
	BidanID    int                     `gorm:"column:bidan_id;not null"`
	Bidan      bidanEntity.Bidan       `gorm:"foreignKey:bidan_id;references:id"`
	PosyanduID int                     `gorm:"column:posyandu_id;not null"`
	Posyandu   posyanduEntity.Posyandu `gorm:"foreignKey:posyandu_id;references:id"`
	CreatedAt  time.Time               `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time               `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Pengampu) TableName() string {
	return "pengampu_posyandu"
}
