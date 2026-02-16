package fakes

import (
	aport "clirzy/auth/internal/application/port"
	"context"
)

type FakeUserGateway struct {
	User *aport.UserDTO
	Err  error
}

func (f *FakeUserGateway) GetByEmail(ctx context.Context, email string) (*aport.UserDTO, error) {
	return f.User, f.Err
}

func (f *FakeUserGateway) RegisterUser(ctx context.Context, email string, password string) (*aport.UserDTO, error) {
	return f.User, f.Err
}
