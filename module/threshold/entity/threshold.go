package entity

type Threshold struct {
	Parameter string  `gorm:"column:parameter;primary_key"`
	Threshold float64 `gorm:"column:threshold;not null"`
}

func (Threshold) TableName() string {
	return "pemeriksaan_threshold"
}
