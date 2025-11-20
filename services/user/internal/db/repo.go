package db

import (
	"context"

	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) FindOneById(ctx context.Context, id uint) (*User, error) {
	user, err := gorm.G[User](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepo) CreateOne(ctx context.Context, user *User) error {
	err := gorm.G[User](r.db).Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) CreateMany(ctx context.Context, users *[]User) error {
	if err := r.db.WithContext(ctx).Create(users).Error; err != nil {
		return err
	}

	return nil
}
