package model

import "time"

type User struct {
	ID           string `gorm:"primaryKey;type:uuid"`
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Status       string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
