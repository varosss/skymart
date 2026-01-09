package repo

import (
	"clirzy/auth/internal/domain/entity"
	"clirzy/auth/internal/domain/port"
	"clirzy/auth/internal/domain/valueobject"
	"clirzy/auth/internal/infrastructure/adapter/gorm/mapper"
	"clirzy/auth/internal/infrastructure/adapter/gorm/model"
	"context"

	"gorm.io/gorm"
)

type RefreshTokensGormRepo struct {
	db *gorm.DB
}

func NewRefreshTokensGormRepo(db *gorm.DB) port.RefreshTokensRepo {
	return &RefreshTokensGormRepo{db: db}
}

func (r *RefreshTokensGormRepo) Save(
	ctx context.Context,
	token *entity.RefreshToken,
) error {
	m := mapper.ToModel(token)

	return r.db.WithContext(ctx).
		Save(m).
		Error
}

func (r *RefreshTokensGormRepo) Get(
	ctx context.Context,
	id valueobject.TokenID,
) (*entity.RefreshToken, error) {
	var m model.RefreshToken

	err := r.db.WithContext(ctx).
		First(&m, "id = ?", id.String()).
		Error

	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return nil, nil
	// }

	if err != nil {
		return nil, err
	}

	return mapper.ToEntity(&m), nil
}

func (r *RefreshTokensGormRepo) Revoke(
	ctx context.Context,
	id valueobject.TokenID,
) error {
	return r.db.WithContext(ctx).
		Model(&model.RefreshToken{}).
		Where("id = ?", id.String()).
		Update("revoked", true).
		Error
}
