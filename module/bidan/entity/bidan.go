package entity

import (
	"github.com/fathoor/posyandu-api/module/user/entity"
	"time"
)

type Bidan struct {
	ID        int         `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    int         `gorm:"column:user_id;not null"`
	User      entity.User `gorm:"foreignKey:user_id;references:id"`
	Jabatan   string      `gorm:"column:jabatan;not null"`
	CreatedAt time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Bidan) TableName() string {
	return "bidan"
}
