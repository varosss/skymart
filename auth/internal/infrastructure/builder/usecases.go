package builder

import (
	aport "clirzy/auth/internal/application/port"
	"clirzy/auth/internal/application/usecase"
	"clirzy/auth/internal/domain/port"
	"time"
)

type UseCases struct {
	Register            *usecase.RegisterUseCase
	Login               *usecase.LoginUseCase
	Logout              *usecase.LogoutUseCase
	RefreshToken        *usecase.RefreshTokenUseCase
	ValidateAccessToken *usecase.ValidateAccessTokenUseCase
}

func BuildUseCases(
	users aport.UserGateway,
	passwords port.PasswordVerifier,
	refreshTokens port.RefreshTokensRepo,
	signer port.TokenSigner,
	verifier port.TokenVerifier,
	clock port.Clock,
	refreshTTL time.Duration,
) *UseCases {
	return &UseCases{
		Register: usecase.NewRegisterUseCase(users),
		Login: usecase.NewLoginUseCase(
			users,
			passwords,
			refreshTokens,
			signer,
			clock,
			refreshTTL,
		),
		Logout: usecase.NewLogoutUseCase(
			verifier,
			refreshTokens,
		),
		RefreshToken: usecase.NewRefreshTokenUseCase(
			refreshTokens,
			verifier,
			signer,
			clock,
			refreshTTL,
		),
		ValidateAccessToken: usecase.NewValidateAccessTokenUseCase(
			verifier,
		),
	}
}
