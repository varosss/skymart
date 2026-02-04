package port

import (
	"context"
)

type UserDTO struct {
	ID string
}

type UserGateway interface {
	FindByID(ctx context.Context, id string) (*UserDTO, error)
}
