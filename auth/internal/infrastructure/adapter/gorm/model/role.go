package model

type Role struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Code string `gorm:"uniqueIndex;not null"`
}
