package port

import (
	"clirzy/user/internal/domain/entity"
	"clirzy/user/internal/domain/valueobject"
	"context"
)

type UsersRepo interface {
	Save(ctx context.Context, user *entity.User) error
	SaveMany(ctx context.Context, users []*entity.User) error

	FindByID(ctx context.Context, id valueobject.UserID) (*entity.User, error)
	FindByEmail(ctx context.Context, email valueobject.Email) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)

	ExistsByEmail(ctx context.Context, email string) bool

	Delete(ctx context.Context, id valueobject.UserID) error
}
