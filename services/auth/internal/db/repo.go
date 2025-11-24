package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type TokensRepo struct {
	db *gorm.DB
}

func NewTokensRepo(db *gorm.DB) *TokensRepo {
	return &TokensRepo{db: db}
}

func (r *TokensRepo) Save(ctx context.Context, userId int, token string, expires time.Time) error {
	err := gorm.G[Token](r.db).Create(
		ctx,
		&Token{
			RefreshToken: token,
			UserId:       uint(userId),
			ExpiresAt:    expires,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *TokensRepo) Find(ctx context.Context, token string) (*Token, error) {
	user, err := gorm.G[Token](r.db).Where("refresh_token = ?", token).First(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *TokensRepo) Delete(ctx context.Context, token string) error {
	_, err := gorm.G[Token](r.db).Where("refresh_token = ?", token).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
