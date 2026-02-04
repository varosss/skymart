package port

import (
	"clirzy/user/internal/domain/entity"
	"clirzy/user/internal/domain/valueobject"
	"context"
)

type UserRepo interface {
	Save(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id valueobject.UserID) (*entity.User, error)
	FindByEmail(ctx context.Context, email valueobject.Email) (*entity.User, error)
	ExistsByEmail(ctx context.Context, email valueobject.Email) bool
}
