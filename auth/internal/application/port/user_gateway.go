package port

import (
	"clirzy/auth/internal/domain/valueobject"
	"context"
)

type UserDTO struct {
	ID           valueobject.UserID
	Email        string
	PasswordHash string
}

type UserGateway interface {
	GetByEmail(ctx context.Context, email string) (*UserDTO, error)
	RegisterUser(ctx context.Context, email string, password string) (*UserDTO, error)
}
