package port

import (
	"clirzy/auth/internal/domain/entity"
	"clirzy/auth/internal/domain/valueobject"
	"context"
)

type RefreshTokensRepo interface {
	Save(ctx context.Context, token *entity.RefreshToken) error
	Get(ctx context.Context, id valueobject.TokenID) (*entity.RefreshToken, error)
	Revoke(ctx context.Context, id valueobject.TokenID) error
}
