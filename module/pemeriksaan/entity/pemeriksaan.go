package entity

import (
	"github.com/fathoor/posyandu-api/module/posyandu/entity"
	remajaEntity "github.com/fathoor/posyandu-api/module/remaja/entity"
	"time"
)

type Pemeriksaan struct {
	ID              int                 `gorm:"column:id;primary_key;auto_increment"`
	PosyanduID      int                 `gorm:"column:posyandu_id;not null"`
	Posyandu        entity.Posyandu     `gorm:"foreignkey:posyandu_id;references:id"`
	RemajaID        int                 `gorm:"column:remaja_id;not null"`
	Remaja          remajaEntity.Remaja `gorm:"foreignkey:remaja_id;references:id"`
	BeratBadan      float64             `gorm:"column:berat_badan"`
	TinggiBadan     float64             `gorm:"column:tinggi_badan"`
	Sistole         float64             `gorm:"column:sistole"`
	Diastole        float64             `gorm:"column:diastole"`
	LingkarLengan   float64             `gorm:"column:lingkar_lengan"`
	TingkatGlukosa  float64             `gorm:"column:tingkat_glukosa"`
	KadarHemoglobin float64             `gorm:"column:kadar_hemoglobin"`
	PemberianFe     bool                `gorm:"column:pemberian_fe;default:false"`
	WaktuPengukuran time.Time           `gorm:"column:waktu_pengukuran;not null"`
	KondisiUmum     string              `gorm:"column:kondisi_umum;not null"`
	Keterangan      string              `gorm:"column:keterangan"`
	CreatedAt       time.Time           `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time           `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Pemeriksaan) TableName() string {
	return "pemeriksaan"
}
