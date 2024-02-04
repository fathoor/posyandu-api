package service

import "github.com/fathoor/posyandu-api/module/chat/model"

type ChatService interface {
	Create(id int, request *model.ChatCreateRequest) (model.ChatResponse, error)
	CreateRoom(id int, request *model.ChatRoomCreateRequest) (model.ChatRoomResponse, error)
	GetByRoomID(id int, roomID string) ([]model.ChatResponse, error)
	GetBySenderID(id, senderID int) ([]model.ChatResponse, error)
	Update(userID int, id string, request *model.ChatUpdateRequest) (model.ChatResponse, error)
	Delete(userID int, id string) error
}
