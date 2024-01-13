package model

type BidanCreateRequest struct {
	UserID  int    `json:"user_id" validate:"required"`
	Jabatan string `json:"jabatan" validate:"required"`
}

type BidanUpdateRequest struct {
	Jabatan string `json:"jabatan" validate:"required"`
}

type BidanUserResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	NIK          int64  `json:"nik"`
	TanggalLahir string `json:"tanggal_lahir"`
	Foto         string `json:"foto"`
	Role         string `json:"role"`
}

type BidanResponse struct {
	ID      int               `json:"id"`
	User    BidanUserResponse `json:"user"`
	Jabatan string            `json:"jabatan"`
}
