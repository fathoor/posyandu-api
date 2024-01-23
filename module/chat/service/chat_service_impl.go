package service

import (
	"github.com/google/uuid"
	"github.com/itsLeonB/posyandu-api/core/exception"
	"github.com/itsLeonB/posyandu-api/module/chat/entity"
	"github.com/itsLeonB/posyandu-api/module/chat/model"
	chatRepository "github.com/itsLeonB/posyandu-api/module/chat/repository"
	"github.com/itsLeonB/posyandu-api/module/chat/validation"
	userRepository "github.com/itsLeonB/posyandu-api/module/user/repository"
)

type chatServiceImpl struct {
	chatRepo chatRepository.ChatRepository
	userRepo userRepository.UserRepository
}

func (service *chatServiceImpl) Create(id int, request *model.ChatCreateRequest) (model.ChatResponse, error) {
	valid := validation.ValidateChatCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	if id != request.SenderID {
		panic(exception.UnauthorizedError{
			Message: "Unauthorized",
		})
	}

	chat := entity.Chat{
		ID:         uuid.New().String(),
		RoomID:     request.RoomID,
		SenderID:   request.SenderID,
		ReceiverID: request.ReceiverID,
		Message:    request.Message,
	}

	err := service.chatRepo.Insert(&chat)
	exception.PanicIfNeeded(err)

	room, err := service.chatRepo.FindRoomByID(chat.RoomID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Room not found",
		})
	}

	konsultan, err := service.userRepo.FindByID(room.KonsultanID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Konsultan not found",
		})
	}

	pasien, err := service.userRepo.FindByID(room.PasienID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pasien not found",
		})
	}

	sender, err := service.userRepo.FindByID(chat.SenderID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Sender not found",
		})
	}

	receiver, err := service.userRepo.FindByID(chat.ReceiverID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Receiver not found",
		})
	}

	response := model.ChatResponse{
		ID: chat.ID,
		Room: model.ChatRoomResponse{
			ID: room.ID,
			Konsultan: model.UserResponse{
				ID:           konsultan.ID,
				Nama:         konsultan.Nama,
				NIK:          konsultan.NIK,
				TanggalLahir: konsultan.TanggalLahir.Format("2006-01-02"),
				Foto:         konsultan.Foto,
				Role:         konsultan.Role,
			},
			Pasien: model.UserResponse{
				ID:           pasien.ID,
				Nama:         pasien.Nama,
				NIK:          pasien.NIK,
				TanggalLahir: pasien.TanggalLahir.Format("2006-01-02"),
				Foto:         pasien.Foto,
				Role:         pasien.Role,
			},
		},
		Sender: model.UserResponse{
			ID:           sender.ID,
			Nama:         sender.Nama,
			NIK:          sender.NIK,
			TanggalLahir: sender.TanggalLahir.Format("2006-01-02"),
			Foto:         sender.Foto,
			Role:         sender.Role,
		},
		Receiver: model.UserResponse{
			ID:           receiver.ID,
			Nama:         receiver.Nama,
			NIK:          receiver.NIK,
			TanggalLahir: receiver.TanggalLahir.Format("2006-01-02"),
			Foto:         receiver.Foto,
			Role:         receiver.Role,
		},
		Message:   chat.Message,
		IsRead:    chat.IsRead,
		Timestamp: chat.Timestamp.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (service *chatServiceImpl) CreateRoom(id int, request *model.ChatRoomCreateRequest) (model.ChatRoomResponse, error) {
	valid := validation.ValidateChatRoomCreateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	if id != request.PasienID {
		panic(exception.UnauthorizedError{
			Message: "Unauthorized",
		})
	}

	room := entity.Room{
		ID:          uuid.New().String(),
		KonsultanID: request.KonsultanID,
		PasienID:    request.PasienID,
	}

	err := service.chatRepo.InsertRoom(&room)
	exception.PanicIfNeeded(err)

	konsultan, err := service.userRepo.FindByID(room.KonsultanID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Konsultan not found",
		})
	}

	pasien, err := service.userRepo.FindByID(room.PasienID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pasien not found",
		})
	}

	response := model.ChatRoomResponse{
		ID: room.ID,
		Konsultan: model.UserResponse{
			ID:           konsultan.ID,
			Nama:         konsultan.Nama,
			NIK:          konsultan.NIK,
			TanggalLahir: konsultan.TanggalLahir.Format("2006-01-02"),
			Foto:         konsultan.Foto,
			Role:         konsultan.Role,
		},
		Pasien: model.UserResponse{
			ID:           pasien.ID,
			Nama:         pasien.Nama,
			NIK:          pasien.NIK,
			TanggalLahir: pasien.TanggalLahir.Format("2006-01-02"),
			Foto:         pasien.Foto,
			Role:         pasien.Role,
		},
	}

	return response, nil
}

func (service *chatServiceImpl) GetByRoomID(id int, roomID string) ([]model.ChatResponse, error) {
	chat, err := service.chatRepo.FindByRoomID(roomID)
	exception.PanicIfNeeded(err)

	if id != chat[0].SenderID && id != chat[0].ReceiverID {
		panic(exception.UnauthorizedError{
			Message: "Unauthorized",
		})
	}

	response := make([]model.ChatResponse, len(chat))
	for i, chat := range chat {
		room, err := service.chatRepo.FindRoomByID(chat.RoomID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Room not found",
			})
		}

		konsultan, err := service.userRepo.FindByID(room.KonsultanID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Konsultan not found",
			})
		}

		pasien, err := service.userRepo.FindByID(room.PasienID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Pasien not found",
			})
		}

		sender, err := service.userRepo.FindByID(chat.SenderID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Sender not found",
			})
		}

		receiver, err := service.userRepo.FindByID(chat.ReceiverID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Receiver not found",
			})
		}

		response[i] = model.ChatResponse{
			ID: chat.ID,
			Room: model.ChatRoomResponse{
				ID: room.ID,
				Konsultan: model.UserResponse{
					ID:           konsultan.ID,
					Nama:         konsultan.Nama,
					NIK:          konsultan.NIK,
					TanggalLahir: konsultan.TanggalLahir.Format("2006-01-02"),
					Foto:         konsultan.Foto,
					Role:         konsultan.Role,
				},
				Pasien: model.UserResponse{
					ID:           pasien.ID,
					Nama:         pasien.Nama,
					NIK:          pasien.NIK,
					TanggalLahir: pasien.TanggalLahir.Format("2006-01-02"),
					Foto:         pasien.Foto,
					Role:         pasien.Role,
				},
			},
			Sender: model.UserResponse{
				ID:           sender.ID,
				Nama:         sender.Nama,
				NIK:          sender.NIK,
				TanggalLahir: sender.TanggalLahir.Format("2006-01-02"),
				Foto:         sender.Foto,
				Role:         sender.Role,
			},
			Receiver: model.UserResponse{
				ID:           receiver.ID,
				Nama:         receiver.Nama,
				NIK:          receiver.NIK,
				TanggalLahir: receiver.TanggalLahir.Format("2006-01-02"),
				Foto:         receiver.Foto,
				Role:         receiver.Role,
			},
			Message:   chat.Message,
			IsRead:    chat.IsRead,
			Timestamp: chat.Timestamp.Format("2006-01-02 15:04:05"),
		}
	}

	return response, nil
}

func (service *chatServiceImpl) GetBySenderID(id, senderID int) ([]model.ChatResponse, error) {
	if id != senderID {
		panic(exception.UnauthorizedError{
			Message: "Unauthorized",
		})
	}

	chat, err := service.chatRepo.FindBySenderID(senderID)
	exception.PanicIfNeeded(err)

	response := make([]model.ChatResponse, len(chat))
	for i, chat := range chat {
		room, err := service.chatRepo.FindRoomByID(chat.RoomID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Room not found",
			})
		}

		konsultan, err := service.userRepo.FindByID(room.KonsultanID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Konsultan not found",
			})
		}

		pasien, err := service.userRepo.FindByID(room.PasienID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Pasien not found",
			})
		}

		sender, err := service.userRepo.FindByID(chat.SenderID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Sender not found",
			})
		}

		receiver, err := service.userRepo.FindByID(chat.ReceiverID)
		if err != nil {
			panic(exception.NotFoundError{
				Message: "Receiver not found",
			})
		}

		response[i] = model.ChatResponse{
			ID: chat.ID,
			Room: model.ChatRoomResponse{
				ID: room.ID,
				Konsultan: model.UserResponse{
					ID:           konsultan.ID,
					Nama:         konsultan.Nama,
					NIK:          konsultan.NIK,
					TanggalLahir: konsultan.TanggalLahir.Format("2006-01-02"),
					Foto:         konsultan.Foto,
					Role:         konsultan.Role,
				},
				Pasien: model.UserResponse{
					ID:           pasien.ID,
					Nama:         pasien.Nama,
					NIK:          pasien.NIK,
					TanggalLahir: pasien.TanggalLahir.Format("2006-01-02"),
					Foto:         pasien.Foto,
					Role:         pasien.Role,
				},
			},
			Sender: model.UserResponse{
				ID:           sender.ID,
				Nama:         sender.Nama,
				NIK:          sender.NIK,
				TanggalLahir: sender.TanggalLahir.Format("2006-01-02"),
				Foto:         sender.Foto,
				Role:         sender.Role,
			},
			Receiver: model.UserResponse{
				ID:           receiver.ID,
				Nama:         receiver.Nama,
				NIK:          receiver.NIK,
				TanggalLahir: receiver.TanggalLahir.Format("2006-01-02"),
				Foto:         receiver.Foto,
				Role:         receiver.Role,
			},
			Message:   chat.Message,
			IsRead:    chat.IsRead,
			Timestamp: chat.Timestamp.Format("2006-01-02 15:04:05"),
		}
	}

	return response, nil
}

func (service *chatServiceImpl) Update(userID int, id string, request *model.ChatUpdateRequest) (model.ChatResponse, error) {
	valid := validation.ValidateChatUpdateRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	chat, err := service.chatRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Chat not found",
		})
	}

	if userID != chat.SenderID && userID != chat.ReceiverID {
		panic(exception.UnauthorizedError{
			Message: "Unauthorized",
		})
	}

	if chat != (entity.Chat{}) {
		chat.RoomID = request.RoomID
		chat.SenderID = request.SenderID
		chat.ReceiverID = request.ReceiverID
		chat.IsRead = request.IsRead
	}

	err = service.chatRepo.Update(&chat)
	exception.PanicIfNeeded(err)

	room, err := service.chatRepo.FindRoomByID(chat.RoomID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Room not found",
		})
	}

	konsultan, err := service.userRepo.FindByID(room.KonsultanID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Konsultan not found",
		})
	}

	pasien, err := service.userRepo.FindByID(room.PasienID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pasien not found",
		})
	}

	sender, err := service.userRepo.FindByID(chat.SenderID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Sender not found",
		})
	}

	receiver, err := service.userRepo.FindByID(chat.ReceiverID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Receiver not found",
		})
	}

	response := model.ChatResponse{
		ID: chat.ID,
		Room: model.ChatRoomResponse{
			ID: room.ID,
			Konsultan: model.UserResponse{
				ID:           konsultan.ID,
				Nama:         konsultan.Nama,
				NIK:          konsultan.NIK,
				TanggalLahir: konsultan.TanggalLahir.Format("2006-01-02"),
				Foto:         konsultan.Foto,
				Role:         konsultan.Role,
			},
			Pasien: model.UserResponse{
				ID:           pasien.ID,
				Nama:         pasien.Nama,
				NIK:          pasien.NIK,
				TanggalLahir: pasien.TanggalLahir.Format("2006-01-02"),
				Foto:         pasien.Foto,
				Role:         pasien.Role,
			},
		},
		Sender: model.UserResponse{
			ID:           sender.ID,
			Nama:         sender.Nama,
			NIK:          sender.NIK,
			TanggalLahir: sender.TanggalLahir.Format("2006-01-02"),
			Foto:         sender.Foto,
			Role:         sender.Role,
		},
		Receiver: model.UserResponse{
			ID:           receiver.ID,
			Nama:         receiver.Nama,
			NIK:          receiver.NIK,
			TanggalLahir: receiver.TanggalLahir.Format("2006-01-02"),
			Foto:         receiver.Foto,
			Role:         receiver.Role,
		},
		Message:   chat.Message,
		IsRead:    chat.IsRead,
		Timestamp: chat.Timestamp.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (service *chatServiceImpl) Delete(userID int, id string) error {
	chat, err := service.chatRepo.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Chat not found",
		})
	}

	if userID != chat.SenderID && userID != chat.ReceiverID {
		panic(exception.UnauthorizedError{
			Message: "Unauthorized",
		})
	}

	return service.chatRepo.Delete(&chat)
}

func ProvideChatService(chatRepo *chatRepository.ChatRepository, userRepo *userRepository.UserRepository) ChatService {
	return &chatServiceImpl{*chatRepo, *userRepo}
}
