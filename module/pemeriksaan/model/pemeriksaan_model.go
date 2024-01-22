package model

import "time"

type PemeriksaanCreateRequest struct {
	PosyanduID      int       `json:"posyandu_id" validate:"required"`
	RemajaID        int       `json:"remaja_id" validate:"required"`
	BeratBadan      float64   `json:"berat_badan"`
	TinggiBadan     float64   `json:"tinggi_badan"`
	Sistole         float64   `json:"sistole"`
	Diastole        float64   `json:"diastole"`
	LingkarLengan   float64   `json:"lingkar_lengan"`
	TingkatGlukosa  float64   `json:"tingkat_glukosa"`
	KadarHemoglobin float64   `json:"kadar_hemoglobin"`
	PemberianFe     bool      `json:"pemberian_fe"`
	WaktuPengukuran time.Time `json:"waktu_pengukuran" validate:"required"`
	KondisiUmum     string    `json:"kondisi_umum" validate:"required"`
	Keterangan      string    `json:"keterangan"`
}

type PemeriksaanUpdateRequest struct {
	PosyanduID      int       `json:"posyandu_id" validate:"required"`
	RemajaID        int       `json:"remaja_id" validate:"required"`
	BeratBadan      float64   `json:"berat_badan"`
	TinggiBadan     float64   `json:"tinggi_badan"`
	Sistole         float64   `json:"sistole"`
	Diastole        float64   `json:"diastole"`
	LingkarLengan   float64   `json:"lingkar_lengan"`
	TingkatGlukosa  float64   `json:"tingkat_glukosa"`
	KadarHemoglobin float64   `json:"kadar_hemoglobin"`
	PemberianFe     bool      `json:"pemberian_fe"`
	WaktuPengukuran time.Time `json:"waktu_pengukuran" validate:"required"`
	KondisiUmum     string    `json:"kondisi_umum" validate:"required"`
	Keterangan      string    `json:"keterangan"`
}

type PemeriksaanPosyanduResponse struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto   string `json:"foto"`
}

type PemeriksaanUserResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	NIK          int64  `json:"nik"`
	TanggalLahir string `json:"tanggal_lahir"`
	Foto         string `json:"foto"`
	Role         string `json:"role"`
}

type PemeriksaanRemajaResponse struct {
	ID       int                         `json:"id"`
	Posyandu PemeriksaanPosyanduResponse `json:"posyandu"`
	User     PemeriksaanUserResponse     `json:"user"`
	NamaAyah string                      `json:"nama_ayah"`
	NamaIbu  string                      `json:"nama_ibu"`
	IsKader  bool                        `json:"is_kader"`
}

type PemeriksaanResponse struct {
	ID              int                         `json:"id"`
	Posyandu        PemeriksaanPosyanduResponse `json:"posyandu"`
	Remaja          PemeriksaanRemajaResponse   `json:"remaja"`
	BeratBadan      float64                     `json:"berat_badan"`
	TinggiBadan     float64                     `json:"tinggi_badan"`
	Sistole         float64                     `json:"sistole"`
	Diastole        float64                     `json:"diastole"`
	LingkarLengan   float64                     `json:"lingkar_lengan"`
	TingkatGlukosa  float64                     `json:"tingkat_glukosa"`
	KadarHemoglobin float64                     `json:"kadar_hemoglobin"`
	PemberianFe     bool                        `json:"pemberian_fe"`
	WaktuPengukuran string                      `json:"waktu_pengukuran"`
	KondisiUmum     string                      `json:"kondisi_umum"`
	Keterangan      string                      `json:"keterangan"`
}
