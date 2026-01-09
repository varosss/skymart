package mapper

import (
	"clirzy/auth/internal/domain/entity"
	"clirzy/auth/internal/domain/valueobject"
	"clirzy/auth/internal/infrastructure/adapter/gorm/model"
)

func ToModel(e *entity.RefreshToken) *model.RefreshToken {
	return &model.RefreshToken{
		ID:        e.ID.String(),
		UserID:    e.UserID.String(),
		ExpiresAt: e.ExpiresAt,
		Revoked:   e.Revoked,
	}
}

func ToEntity(m *model.RefreshToken) *entity.RefreshToken {
	return &entity.RefreshToken{
		ID:        valueobject.TokenID(m.ID),
		UserID:    valueobject.UserID(m.UserID),
		ExpiresAt: m.ExpiresAt,
		Revoked:   m.Revoked,
	}
}
