package model

type PengampuCreateRequest struct {
	BidanID    int  `json:"bidan_id" validate:"required"`
	PosyanduID int  `json:"posyandu_id" validate:"required"`
	Active     bool `json:"active" validate:"required"`
}

type PengampuUpdateRequest struct {
	BidanID    int  `json:"bidan_id" validate:"required"`
	PosyanduID int  `json:"posyandu_id" validate:"required"`
	Active     bool `json:"active" validate:"required"`
}

type PengampuBidanUserResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	NIK          int64  `json:"nik"`
	TanggalLahir string `json:"tanggal_lahir"`
	Foto         string `json:"foto"`
	Role         string `json:"role"`
}

type PengampuBidanResponse struct {
	ID      int                       `json:"id"`
	User    PengampuBidanUserResponse `json:"user"`
	Jabatan string                    `json:"jabatan"`
}

type PengampuPosyanduResponse struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto   string `json:"foto"`
}

type PengampuResponse struct {
	Bidan    PengampuBidanResponse    `json:"bidan"`
	Posyandu PengampuPosyanduResponse `json:"posyandu"`
	Active   bool                     `json:"active"`
}
