package port

import (
	"context"
)

type UserDTO struct {
	ID string
}

type UserQuery interface {
	FindByID(ctx context.Context, id string) (*UserDTO, error)
}
