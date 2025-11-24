package service

import (
	"context"
	"errors"
	"time"

	userclient "clirzy/pkg/client/user"
	"clirzy/pkg/consts"
	"clirzy/pkg/security"
	"clirzy/pkg/utils"
	"clirzy/services/auth/internal/db"
	"clirzy/services/auth/internal/domain"
)

type AuthService struct {
	repo       *db.TokensRepo
	jwtManager *utils.JWTManager
	userClient *userclient.Client
}

func NewAuthService(
	repo *db.TokensRepo,
	jwtManager *utils.JWTManager,
	userClient *userclient.Client,
) *AuthService {
	return &AuthService{
		repo:       repo,
		jwtManager: jwtManager,
		userClient: userClient,
	}
}

func (s *AuthService) Register(ctx context.Context, input domain.AuthInput) (*domain.User, error) {
	createResp, err := s.userClient.CreateUser(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	user := createResp.Users[0]

	return &domain.User{Id: uint(user.Id), Username: user.Username}, nil
}

func (s *AuthService) Login(ctx context.Context, input domain.AuthInput) (*domain.AuthPair, error) {
	user, err := s.userClient.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return nil, err
	}

	if err = security.ComparePassword(user.PasswordHash, input.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	newAccess, err := s.generateAccessToken(int(user.Id))
	if err != nil {
		return nil, errors.New("couldn't generate auth token")
	}

	newRefresh, err := s.generateRefreshToken()
	if err != nil {
		return nil, errors.New("couldn't generate refresh token")
	}

	err = s.repo.Save(ctx, int(user.Id), newRefresh, time.Now().Add(consts.DEFAULT_REFRESH_TOKEN_LIFETIME))

	if err != nil {
		return nil, errors.New("couldn't create refresh token in db")
	}

	return &domain.AuthPair{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}, nil
}

func (s *AuthService) Refresh(ctx context.Context, token string) (*domain.AuthPair, error) {
	rt, err := s.repo.Find(ctx, token)
	if err != nil {
		return nil, errors.New("refresh token is invalid")
	}

	if time.Now().After(rt.ExpiresAt) {
		return nil, errors.New("refresh token is expired")
	}

	user, err := s.userClient.GetUserById(ctx, rt.UserId)
	if err != nil {
		return nil, err
	}

	newAccess, err := s.generateAccessToken(int(user.Id))
	if err != nil {
		return nil, err
	}

	newRefresh, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	_ = s.repo.Delete(ctx, token)

	err = s.repo.Save(ctx, int(user.Id), newRefresh, time.Now().Add(consts.DEFAULT_REFRESH_TOKEN_LIFETIME))
	if err != nil {
		return nil, err
	}

	return &domain.AuthPair{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}, nil
}

func (s *AuthService) generateAccessToken(userId int) (string, error) {
	accessToken, err := s.jwtManager.Generate(int64(userId))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *AuthService) generateRefreshToken() (string, error) {
	refreshToken, err := security.GenerateRefreshToken()
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
