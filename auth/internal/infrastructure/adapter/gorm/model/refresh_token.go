package model

import "time"

type RefreshToken struct {
	ID        string    `gorm:"primaryKey;size:64"`
	UserID    string    `gorm:"index;size:64;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Revoked   bool      `gorm:"not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
