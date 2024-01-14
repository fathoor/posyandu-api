package model

import "time"

type JadwalPosyanduCreateRequest struct {
	PosyanduID   int       `json:"posyandu_id" validate:"required"`
	WaktuMulai   time.Time `json:"waktu_mulai" validate:"required"`
	WaktuSelesai time.Time `json:"waktu_selesai" validate:"required"`
}

type JadwalPosyanduUpdateRequest struct {
	PosyanduID   int       `json:"posyandu_id" validate:"required"`
	WaktuMulai   time.Time `json:"waktu_mulai" validate:"required"`
	WaktuSelesai time.Time `json:"waktu_selesai" validate:"required"`
}

type JadwalPosyanduPosyanduResponse struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto   string `json:"foto"`
}

type JadwalPosyanduResponse struct {
	ID           int                            `json:"id"`
	Posyandu     JadwalPosyanduPosyanduResponse `json:"posyandu"`
	WaktuMulai   string                         `json:"waktu_mulai"`
	WaktuSelesai string                         `json:"waktu_selesai"`
}
