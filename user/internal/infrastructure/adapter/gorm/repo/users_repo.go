// internal/infrastructure/adapter/gorm/repo/users_repo.go
package repo

import (
	"clirzy/user/internal/domain/entity"
	"clirzy/user/internal/domain/valueobject"
	"clirzy/user/internal/infrastructure/adapter/gorm/mapper"
	"clirzy/user/internal/infrastructure/adapter/gorm/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UsersGormRepo struct {
	db *gorm.DB
}

func NewUsersGormRepo(db *gorm.DB) *UsersGormRepo {
	return &UsersGormRepo{db: db}
}

func (r *UsersGormRepo) Save(
	ctx context.Context,
	user *entity.User,
) error {

	m := mapper.ToModel(user)

	return r.db.WithContext(ctx).
		Save(m).
		Error
}

func (r *UsersGormRepo) FindByID(
	ctx context.Context,
	id valueobject.UserID,
) (*entity.User, error) {

	var m model.User

	err := r.db.WithContext(ctx).
		First(&m, "id = ?", id.String()).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return mapper.ToDomain(&m)
}

func (r *UsersGormRepo) FindByEmail(
	ctx context.Context,
	email valueobject.Email,
) (*entity.User, error) {

	var m model.User

	err := r.db.WithContext(ctx).
		First(&m, "email = ?", email.String()).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return mapper.ToDomain(&m)
}

func (r *UsersGormRepo) ExistsByEmail(
	ctx context.Context,
	email valueobject.Email,
) bool {

	var count int64

	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Where("email = ?", email.String()).
		Count(&count).
		Error

	if err != nil {
		return false
	}

	return count > 0
}
