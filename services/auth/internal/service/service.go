package service

import (
	"context"
	"errors"
)

// AuthService реализует всю логику авторизации
type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	if email == "test@test.com" && password == "123" {
		return "hardcoded_token", nil
	}
	return "", errors.New("invalid credentials")
}
