package model

import "time"

type JadwalPenyuluhanCreateRequest struct {
	PosyanduID   int       `json:"posyandu_id" validate:"required"`
	WaktuMulai   time.Time `json:"waktu_mulai" validate:"required"`
	WaktuSelesai time.Time `json:"waktu_selesai" validate:"required"`
	Title        string    `json:"title" validate:"required"`
	Materi       string    `json:"materi"`
	Feedback     string    `json:"feedback"`
}

type JadwalPenyuluhanUpdateRequest struct {
	PosyanduID   int       `json:"posyandu_id" validate:"required"`
	WaktuMulai   time.Time `json:"waktu_mulai" validate:"required"`
	WaktuSelesai time.Time `json:"waktu_selesai" validate:"required"`
	Title        string    `json:"title" validate:"required"`
	Materi       string    `json:"materi"`
	Feedback     string    `json:"feedback"`
}

type JadwalPenyuluhanPosyanduResponse struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto   string `json:"foto"`
}

type JadwalPenyuluhanResponse struct {
	ID           int                              `json:"id"`
	Posyandu     JadwalPenyuluhanPosyanduResponse `json:"posyandu"`
	WaktuMulai   string                           `json:"waktu_mulai"`
	WaktuSelesai string                           `json:"waktu_selesai"`
	Title        string                           `json:"title"`
	Materi       string                           `json:"materi"`
	Feedback     string                           `json:"feedback"`
}
