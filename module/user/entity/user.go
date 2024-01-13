package entity

import "time"

type User struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement"`
	Nama         string    `gorm:"column:nama;not null"`
	Email        string    `gorm:"column:email;not null;unique"`
	Username     string    `gorm:"column:username;not null;unique"`
	Password     string    `gorm:"column:password;not null"`
	NIK          int64     `gorm:"column:nik;not null;unique"`
	TempatLahir  string    `gorm:"column:tempat_lahir;not null"`
	TanggalLahir time.Time `gorm:"column:tanggal_lahir;not null"`
	Alamat       string    `gorm:"column:alamat;not null"`
	Provinsi     string    `gorm:"column:provinsi;not null"`
	Kota         string    `gorm:"column:kota;not null"`
	Kecamatan    string    `gorm:"column:kecamatan;not null"`
	Kelurahan    string    `gorm:"column:kelurahan;not null"`
	KodePos      int       `gorm:"column:kode_pos;not null"`
	RT           int       `gorm:"column:rt;not null"`
	RW           int       `gorm:"column:rw;not null"`
	Telepon      string    `gorm:"column:telepon;not null"`
	Foto         string    `gorm:"column:foto;not null;default:'/storage/image/default.png'"`
	Role         string    `gorm:"column:role;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
