package entity

import (
	"clirzy/auth/internal/domain/valueobject"
	"time"
)

type RefreshToken struct {
	ID        valueobject.TokenID
	UserID    valueobject.UserID
	ExpiresAt time.Time
	Revoked   bool
}

func NewRefreshToken(
	id valueobject.TokenID,
	userID valueobject.UserID,
	expiresAt time.Time,
) *RefreshToken {
	return &RefreshToken{
		ID:        id,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
}

func (t *RefreshToken) IsExpired(now time.Time) bool {
	return now.After(t.ExpiresAt)
}

func (t *RefreshToken) Revoke() {
	t.Revoked = true
}
