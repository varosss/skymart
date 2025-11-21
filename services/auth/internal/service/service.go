package service

import (
	"context"
	"errors"

	userclient "clirzy/pkg/client/user"
	"clirzy/pkg/utils"
	"clirzy/services/auth/internal/domain"
)

type AuthService struct {
	userClient *userclient.Client
}

func NewAuthService(userClient *userclient.Client) *AuthService {
	return &AuthService{userClient: userClient}
}

func (s *AuthService) Register(ctx context.Context, input domain.AuthInput) (*domain.User, error) {
	createResp, err := s.userClient.CreateUser(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	user := createResp.Users[0]

	return &domain.User{Id: uint(user.Id), Username: user.Username}, nil
}

func (s *AuthService) Login(ctx context.Context, input domain.AuthInput) (*domain.LoginOutput, error) {
	user, err := s.userClient.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return nil, err
	}

	if err = utils.ComparePassword(user.PasswordHash, input.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &domain.LoginOutput{AccessToken: "", RefreshToken: ""}, nil
}
