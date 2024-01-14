package model

type HomePosyanduResponse struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto   string `json:"foto"`
}

type HomeUserResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	NIK          int64  `json:"nik"`
	TanggalLahir string `json:"tanggal_lahir"`
	Foto         string `json:"foto"`
	Role         string `json:"role"`
}

type HomeRemajaResponse struct {
	ID       int                  `json:"id"`
	Posyandu HomePosyanduResponse `json:"posyandu"`
	User     HomeUserResponse     `json:"user"`
	NamaAyah string               `json:"nama_ayah"`
	NamaIbu  string               `json:"nama_ibu"`
	IsKader  bool                 `json:"is_kader"`
}

type HomePemeriksaanResponse struct {
	ID              int                `json:"id"`
	Remaja          HomeRemajaResponse `json:"remaja"`
	BeratBadan      float64            `json:"berat_badan"`
	TinggiBadan     float64            `json:"tinggi_badan"`
	TekananDarah    float64            `json:"tekanan_darah"`
	LingkarLengan   float64            `json:"lingkar_lengan"`
	TingkatGlukosa  float64            `json:"tingkat_glukosa"`
	KadarHemoglobin float64            `json:"kadar_hemoglobin"`
	PemberianFe     bool               `json:"pemberian_fe"`
	WaktuPengukuran string             `json:"waktu_pengukuran"`
	KondisiUmum     string             `json:"kondisi_umum"`
	Keterangan      string             `json:"keterangan"`
}

type HomeJadwalPosyanduResponse struct {
	ID           int                  `json:"id"`
	Posyandu     HomePosyanduResponse `json:"posyandu"`
	WaktuMulai   string               `json:"waktu_mulai"`
	WaktuSelesai string               `json:"waktu_selesai"`
}

type HomeJadwalPenyuluhanResponse struct {
	ID           int                  `json:"id"`
	Posyandu     HomePosyanduResponse `json:"posyandu"`
	WaktuMulai   string               `json:"waktu_mulai"`
	WaktuSelesai string               `json:"waktu_selesai"`
	Title        string               `json:"title"`
	Materi       string               `json:"materi"`
	Feedback     string               `json:"feedback"`
}

type BidanHomeResponse struct {
	User             HomeUserResponse               `json:"user"`
	Posyandu         HomePosyanduResponse           `json:"posyandu"`
	Pemeriksaan      []HomePemeriksaanResponse      `json:"pemeriksaan"`
	JadwalPosyandu   []HomeJadwalPosyanduResponse   `json:"jadwal_posyandu"`
	JadwalPenyuluhan []HomeJadwalPenyuluhanResponse `json:"jadwal_penyuluhan"`
}

type HomeResponse struct {
	User             HomeUserResponse               `json:"user"`
	JadwalPosyandu   []HomeJadwalPosyanduResponse   `json:"jadwal_posyandu"`
	JadwalPenyuluhan []HomeJadwalPenyuluhanResponse `json:"jadwal_penyuluhan"`
}
