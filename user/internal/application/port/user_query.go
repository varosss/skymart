package port

import (
	"context"
)

type UserDTO struct {
	ID           string
	Email        string
	PasswordHash string
}

type UserQuery interface {
	GetByEmail(ctx context.Context, email string) (*UserDTO, error)
}
