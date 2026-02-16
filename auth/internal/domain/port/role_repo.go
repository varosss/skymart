package port

import (
	"context"

	"clirzy/auth/internal/domain/entity"
	"clirzy/auth/internal/domain/valueobject"
)

type RoleRepo interface {
	GetByCode(
		ctx context.Context,
		code valueobject.RoleCode,
	) (*entity.Role, error)

	ListAll(
		ctx context.Context,
	) ([]entity.Role, error)
}
