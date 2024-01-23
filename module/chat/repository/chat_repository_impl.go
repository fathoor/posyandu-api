package repository

import (
	"github.com/itsLeonB/posyandu-api/module/chat/entity"
	"gorm.io/gorm"
)

type chatRepositoryImpl struct {
	*gorm.DB
}

func (repository *chatRepositoryImpl) Insert(chat *entity.Chat) error {
	return repository.DB.Create(&chat).Error
}

func (repository *chatRepositoryImpl) InsertRoom(room *entity.Room) error {
	return repository.DB.Create(&room).Error
}

func (repository *chatRepositoryImpl) FindByRoomID(id string) ([]entity.Chat, error) {
	var chat []entity.Chat
	err := repository.DB.Find(&chat, "room_id = ?", id).Error

	return chat, err
}

func (repository *chatRepositoryImpl) FindBySenderID(id int) ([]entity.Chat, error) {
	var chat []entity.Chat
	err := repository.DB.Find(&chat, "sender_id = ? order by timestamp desc", id).Limit(2).Error

	return chat, err
}

func (repository *chatRepositoryImpl) FindByID(id string) (entity.Chat, error) {
	var chat entity.Chat
	err := repository.DB.Take(&chat, "id = ?", id).Error

	return chat, err
}

func (repository *chatRepositoryImpl) FindRoomByID(id string) (entity.Room, error) {
	var room entity.Room
	err := repository.DB.Take(&room, "id = ?", id).Error

	return room, err
}

func (repository *chatRepositoryImpl) Update(chat *entity.Chat) error {
	return repository.DB.Save(&chat).Error
}

func (repository *chatRepositoryImpl) Delete(chat *entity.Chat) error {
	return repository.DB.Delete(&chat).Error
}

func ProvideChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepositoryImpl{db}
}
