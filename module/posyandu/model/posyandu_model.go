package model

type PosyanduCreateRequest struct {
	Nama      string `json:"nama" validate:"required"`
	Alamat    string `json:"alamat" validate:"required"`
	Provinsi  string `json:"provinsi" validate:"required"`
	Kota      string `json:"kota" validate:"required"`
	Kecamatan string `json:"kecamatan" validate:"required"`
	Kelurahan string `json:"kelurahan" validate:"required"`
	KodePos   int    `json:"kode_pos" validate:"required"`
	RT        int    `json:"rt" validate:"required"`
	RW        int    `json:"rw" validate:"required"`
	Foto      string `json:"foto"`
}

type PosyanduUpdateRequest struct {
	Nama      string `json:"nama" validate:"required"`
	Alamat    string `json:"alamat" validate:"required"`
	Provinsi  string `json:"provinsi" validate:"required"`
	Kota      string `json:"kota" validate:"required"`
	Kecamatan string `json:"kecamatan" validate:"required"`
	Kelurahan string `json:"kelurahan" validate:"required"`
	KodePos   int    `json:"kode_pos" validate:"required"`
	RT        int    `json:"rt" validate:"required"`
	RW        int    `json:"rw" validate:"required"`
	Foto      string `json:"foto"`
}

type PosyanduResponse struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	KodePos   int    `json:"kode_pos"`
	RT        int    `json:"rt"`
	RW        int    `json:"rw"`
	Foto      string `json:"foto"`
}
