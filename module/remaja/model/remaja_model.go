package model

type RemajaCreateRequest struct {
	PosyanduID int    `json:"posyandu_id" validate:"required"`
	UserID     int    `json:"user_id" validate:"required"`
	NamaAyah   string `json:"nama_ayah" validate:"required"`
	NamaIbu    string `json:"nama_ibu" validate:"required"`
}

type RemajaUpdateRequest struct {
	PosyanduID int    `json:"posyandu_id" validate:"required"`
	NamaAyah   string `json:"nama_ayah" validate:"required"`
	NamaIbu    string `json:"nama_ibu" validate:"required"`
}

type RemajaUpdateKaderRequest struct {
	PosyanduID int    `json:"posyandu_id" validate:"required"`
	NamaAyah   string `json:"nama_ayah" validate:"required"`
	NamaIbu    string `json:"nama_ibu" validate:"required"`
	IsKader    bool   `json:"is_kader"`
}

type RemajaPosyanduResponse struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto   string `json:"foto"`
}

type RemajaUserResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	NIK          int64  `json:"nik"`
	TanggalLahir string `json:"tanggal_lahir"`
	Foto         string `json:"foto"`
	Role         string `json:"role"`
}

type RemajaResponse struct {
	ID       int                    `json:"id"`
	Posyandu RemajaPosyanduResponse `json:"posyandu"`
	User     RemajaUserResponse     `json:"user"`
	NamaAyah string                 `json:"nama_ayah"`
	NamaIbu  string                 `json:"nama_ibu"`
	IsKader  bool                   `json:"is_kader"`
}

type PemeriksaanResponse struct {
	ID              int                    `json:"id"`
	Posyandu        RemajaPosyanduResponse `json:"posyandu"`
	Remaja          RemajaResponse         `json:"remaja"`
	BeratBadan      float64                `json:"berat_badan"`
	TinggiBadan     float64                `json:"tinggi_badan"`
	Sistole         float64                `json:"sistole"`
	Diastole        float64                `json:"diastole"`
	LingkarLengan   float64                `json:"lingkar_lengan"`
	TingkatGlukosa  float64                `json:"tingkat_glukosa"`
	KadarHemoglobin float64                `json:"kadar_hemoglobin"`
	PemberianFe     bool                   `json:"pemberian_fe"`
	WaktuPengukuran string                 `json:"waktu_pengukuran"`
	KondisiUmum     string                 `json:"kondisi_umum"`
	Keterangan      string                 `json:"keterangan"`
}

type RemajaPemeriksaanResponse struct {
	ID          int                    `json:"id"`
	Posyandu    RemajaPosyanduResponse `json:"posyandu"`
	User        RemajaUserResponse     `json:"user"`
	NamaAyah    string                 `json:"nama_ayah"`
	NamaIbu     string                 `json:"nama_ibu"`
	IsKader     bool                   `json:"is_kader"`
	Pemeriksaan PemeriksaanResponse    `json:"pemeriksaan"`
}
