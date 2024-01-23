package model

type UserResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	NIK          int64  `json:"nik"`
	TanggalLahir string `json:"tanggal_lahir"`
	Foto         string `json:"foto"`
	Role         string `json:"role"`
}

type ChatRoomCreateRequest struct {
	KonsultanID int `json:"konsultan_id" validate:"required"`
	PasienID    int `json:"pasien_id" validate:"required"`
}

type ChatRoomResponse struct {
	ID        string       `json:"id"`
	Konsultan UserResponse `json:"konsultan"`
	Pasien    UserResponse `json:"pasien"`
}

type ChatCreateRequest struct {
	RoomID     string `json:"room_id" validate:"required"`
	SenderID   int    `json:"sender_id" validate:"required"`
	ReceiverID int    `json:"receiver_id" validate:"required"`
	Message    string `json:"message" validate:"required"`
}

type ChatUpdateRequest struct {
	RoomID     string `json:"room_id" validate:"required"`
	SenderID   int    `json:"sender_id" validate:"required"`
	ReceiverID int    `json:"receiver_id" validate:"required"`
	IsRead     bool   `json:"is_read" validate:"required"`
}

type ChatResponse struct {
	ID        string           `json:"id"`
	Room      ChatRoomResponse `json:"room"`
	Sender    UserResponse     `json:"sender"`
	Receiver  UserResponse     `json:"receiver"`
	Message   string           `json:"message"`
	IsRead    bool             `json:"is_read"`
	Timestamp string           `json:"timestamp"`
}
