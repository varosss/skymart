package query

import (
	aport "clirzy/user/internal/application/port"
	"clirzy/user/internal/domain"
	"clirzy/user/internal/domain/port"
	"clirzy/user/internal/domain/valueobject"
	"context"
)

type UserQuery struct {
	repo port.UserRepo
}

func NewUserQuery(repo port.UserRepo) *UserQuery {
	return &UserQuery{
		repo: repo,
	}
}

func (q *UserQuery) GetByEmail(ctx context.Context, email valueobject.Email) (*aport.UserDTO, error) {
	user, err := q.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return &aport.UserDTO{
		ID:           user.ID().String(),
		Email:        user.Email().String(),
		PasswordHash: user.PasswordHash().String(),
	}, nil
}

func (q *UserQuery) GetByID(ctx context.Context, userID valueobject.UserID) (*aport.UserDTO, error) {
	user, err := q.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return &aport.UserDTO{
		ID:           user.ID().String(),
		Email:        user.Email().String(),
		PasswordHash: user.PasswordHash().String(),
	}, nil
}
