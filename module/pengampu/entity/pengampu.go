package entity

import (
	bidanEntity "github.com/fathoor/posyandu-api/module/bidan/entity"
	posyanduEntity "github.com/fathoor/posyandu-api/module/posyandu/entity"
	"time"
)

type Pengampu struct {
	BidanID    int                     `gorm:"column:bidan_id;primaryKey;not null"`
	Bidan      bidanEntity.Bidan       `gorm:"foreignKey:bidan_id;references:id"`
	PosyanduID int                     `gorm:"column:posyandu_id;primaryKey;not null"`
	Posyandu   posyanduEntity.Posyandu `gorm:"foreignKey:posyandu_id;references:id"`
	Active     bool                    `gorm:"column:active;default:false"`
	CreatedAt  time.Time               `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time               `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Pengampu) TableName() string {
	return "pengampu_posyandu"
}
