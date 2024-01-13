package entity

import (
	posyanduEntity "github.com/itsLeonB/posyandu-api/module/posyandu/entity"
	userEntity "github.com/itsLeonB/posyandu-api/module/user/entity"
	"time"
)

type Remaja struct {
	ID         int                     `gorm:"column:id;primary_key;auto_increment"`
	PosyanduID int                     `gorm:"column:posyandu_id;not null"`
	Posyandu   posyanduEntity.Posyandu `gorm:"foreignkey:posyandu_id;references:id"`
	UserID     int                     `gorm:"column:user_id;not null;unique"`
	User       userEntity.User         `gorm:"foreignkey:user_id;references:id"`
	NamaAyah   string                  `gorm:"column:nama_ayah;not null"`
	NamaIbu    string                  `gorm:"column:nama_ibu;not null"`
	IsKader    bool                    `gorm:"column:is_kader;default:false"`
	CreatedAt  time.Time               `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time               `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Remaja) TableName() string {
	return "remaja"
}
