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

type RefreshTokenGormRepo struct {
	db *gorm.DB
}

func NewRefreshTokenGormRepo(db *gorm.DB) port.RefreshTokenRepo {
	return &RefreshTokenGormRepo{db: db}
}

func (r *RefreshTokenGormRepo) Save(
	ctx context.Context,
	token *entity.RefreshToken,
) error {
	m := mapper.ToModel(token)

	return r.db.WithContext(ctx).
		Save(m).
		Error
}

func (r *RefreshTokenGormRepo) Get(
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

func (r *RefreshTokenGormRepo) Revoke(
	ctx context.Context,
	id valueobject.TokenID,
) error {
	return r.db.WithContext(ctx).
		Model(&model.RefreshToken{}).
		Where("id = ?", id.String()).
		Update("revoked", true).
		Error
}
