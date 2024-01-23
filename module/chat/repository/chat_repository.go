package repository

import "github.com/itsLeonB/posyandu-api/module/chat/entity"

type ChatRepository interface {
	Insert(chat *entity.Chat) error
	InsertRoom(room *entity.Room) error
	FindByRoomID(id string) ([]entity.Chat, error)
	FindBySenderID(id int) ([]entity.Chat, error)
	FindByID(id string) (entity.Chat, error)
	FindRoomByID(id string) (entity.Room, error)
	Update(chat *entity.Chat) error
	Delete(chat *entity.Chat) error
}
