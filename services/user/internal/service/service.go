package service

import (
	"context"

	"clirzy/pkg/security"
	"clirzy/services/user/internal/db"
	"clirzy/services/user/internal/domain"
)

type UserService struct {
	repo *db.UsersRepo
}

func NewUserService(repo *db.UsersRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserById(ctx context.Context, userId int) (*domain.User, error) {
	user, err := s.repo.FindOneById(ctx, uint(userId))
	if err != nil {
		return nil, err
	}

	domainUser := db.DBToDomain(*user)

	return domainUser, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := s.repo.FindOneByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	domainUser := db.DBToDomain(*user)

	return domainUser, nil
}

func (s *UserService) CreateOne(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
	hashed, err := security.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input.Password = hashed

	userModel := db.CreateUserInputToDB(input)

	err = s.repo.CreateOne(ctx, &userModel)
	if err != nil {
		return nil, err
	}

	userDomain := db.DBToDomain(userModel)

	return userDomain, nil
}

func (s *UserService) CreateMany(ctx context.Context, inputs []domain.CreateUserInput) ([]*domain.User, error) {
	userModels := make([]db.User, len(inputs))

	for i, inp := range inputs {
		hashed, err := security.HashPassword(inp.Password)
		if err != nil {
			return nil, err
		}

		inp.Password = hashed

		userModels[i] = db.CreateUserInputToDB(inp)
	}

	err := s.repo.CreateMany(ctx, &userModels)
	if err != nil {
		return nil, err
	}

	userDomains := make([]*domain.User, len(userModels))
	for i, model := range userModels {
		userDomains[i] = db.DBToDomain(model)
	}

	return userDomains, nil
}
