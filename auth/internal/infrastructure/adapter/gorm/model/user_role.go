package model

import (
	"time"
)

type UserRole struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	UserID    string `gorm:"type:uuid;index;not null"`
	RoleID    string `gorm:"type:uuid;index;not null"`
	CreatedAt time.Time
}
