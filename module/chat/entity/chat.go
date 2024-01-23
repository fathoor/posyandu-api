package entity

import (
	"github.com/itsLeonB/posyandu-api/module/user/entity"
	"time"
)

type Chat struct {
	ID         string      `gorm:"column:id;primaryKey"`
	RoomID     string      `gorm:"column:room_id;not null"`
	Room       Room        `gorm:"foreignKey:room_id;references:id"`
	SenderID   int         `gorm:"column:sender_id;not null"`
	Sender     entity.User `gorm:"foreignKey:sender_id;references:id"`
	ReceiverID int         `gorm:"column:receiver_id;not null"`
	Receiver   entity.User `gorm:"foreignKey:receiver_id;references:id"`
	Message    string      `gorm:"column:message;not null"`
	IsRead     bool        `gorm:"column:is_read;default:false"`
	Timestamp  time.Time   `gorm:"column:timestamp;autoCreateTime"`
	CreatedAt  time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Chat) TableName() string {
	return "chat"
}
