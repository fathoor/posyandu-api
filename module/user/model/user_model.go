package model

import "time"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	Role      string `json:"role"`
	ExpiresAt string `json:"expires_at"`
}

type UserRegisterRequest struct {
	Nama         string    `json:"nama" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Username     string    `json:"username" validate:"required"`
	Password     string    `json:"password" validate:"required"`
	NIK          int64     `json:"nik" validate:"required"`
	TempatLahir  string    `json:"tempat_lahir" validate:"required"`
	TanggalLahir time.Time `json:"tanggal_lahir" validate:"required"`
	Alamat       string    `json:"alamat" validate:"required"`
	Provinsi     string    `json:"provinsi" validate:"required"`
	Kota         string    `json:"kota" validate:"required"`
	Kecamatan    string    `json:"kecamatan" validate:"required"`
	Kelurahan    string    `json:"kelurahan" validate:"required"`
	KodePos      int       `json:"kode_pos" validate:"required"`
	RT           int       `json:"rt" validate:"required"`
	RW           int       `json:"rw" validate:"required"`
	Telepon      string    `json:"telepon" validate:"required"`
	Foto         string    `json:"foto"`
	Role         string    `json:"role" validate:"required"`
}

type UserUpdateRequest struct {
	Nama      string `json:"nama" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required"`
	Alamat    string `json:"alamat" validate:"required"`
	Provinsi  string `json:"provinsi" validate:"required"`
	Kota      string `json:"kota" validate:"required"`
	Kecamatan string `json:"kecamatan" validate:"required"`
	Kelurahan string `json:"kelurahan" validate:"required"`
	KodePos   int    `json:"kode_pos" validate:"required"`
	RT        int    `json:"rt" validate:"required"`
	RW        int    `json:"rw" validate:"required"`
	Telepon   string `json:"telepon" validate:"required"`
	Foto      string `json:"foto"`
}

type UserUpdateAuthRequest struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type UserResponse struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	NIK          int64  `json:"nik"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Alamat       string `json:"alamat"`
	Provinsi     string `json:"provinsi"`
	Kota         string `json:"kota"`
	Kecamatan    string `json:"kecamatan"`
	Kelurahan    string `json:"kelurahan"`
	KodePos      int    `json:"kode_pos"`
	RT           int    `json:"rt"`
	RW           int    `json:"rw"`
	Telepon      string `json:"telepon"`
	Foto         string `json:"foto"`
	Role         string `json:"role"`
}
