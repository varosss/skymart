package fakes

import (
	"clirzy/auth/internal/domain/entity"
	"clirzy/auth/internal/domain/valueobject"
	"context"
)

type FakeRefreshTokenRepo struct {
	Token *entity.RefreshToken
	Err   error
}

func (f *FakeRefreshTokenRepo) Save(ctx context.Context, token *entity.RefreshToken) error {
	f.Token = token
	return f.Err
}

func (f *FakeRefreshTokenRepo) Get(ctx context.Context, id valueobject.TokenID) (*entity.RefreshToken, error) {
	return f.Token, f.Err
}

func (f *FakeRefreshTokenRepo) Revoke(ctx context.Context, id valueobject.TokenID) error {
	return f.Err
}
