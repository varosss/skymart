package db

import (
	"time"
)

type Token struct {
	ID           uint   `gorm:"primaryKey"`
	UserId       uint   `gorm:"uniqueIndex"`
	RefreshToken string `gorm:"size:32;not null"`
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
