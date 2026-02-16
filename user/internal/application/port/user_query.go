package port

import (
	"clirzy/user/internal/domain/valueobject"
	"context"
)

type UserDTO struct {
	ID           string
	Email        string
	PasswordHash string
}

type UserQuery interface {
	GetByEmail(ctx context.Context, email valueobject.Email) (*UserDTO, error)
	GetByID(ctx context.Context, userID valueobject.UserID) (*UserDTO, error)
}
