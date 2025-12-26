package port

import (
	"context"
)

type UserDTO struct {
	ID           string
	Email        string
	PasswordHash string
}

type UserService interface {
	FindByEmail(ctx context.Context, email string) (*UserDTO, error)
}
