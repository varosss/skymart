package port

import (
	"context"
)

type UserDTO struct {
	ID           string
	Email        string
	PasswordHash string
}

type UserGateway interface {
	FindByEmail(ctx context.Context, email string) (*UserDTO, error)
	RegisterUser(ctx context.Context, email string, password string) (*UserDTO, error)
}
