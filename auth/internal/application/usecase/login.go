package usecase

import (
	aport "clirzy/auth/internal/application/port"
	"clirzy/auth/internal/domain"
	"clirzy/auth/internal/domain/entity"
	"clirzy/auth/internal/domain/port"
	"clirzy/auth/internal/domain/valueobject"
	"context"
	"time"
)

type LoginCommand struct {
	Email    string
	Password string
}

type LoginResult struct {
	AccessToken  string
	RefreshToken string
}

type LoginUseCase struct {
	users         aport.UserGateway
	passwords     port.PasswordVerifier
	refreshTokens port.RefreshTokensRepo
	signer        port.TokenSigner
	clock         port.Clock
	refreshTTL    time.Duration
}

func NewLoginUseCase(
	users aport.UserGateway,
	passwords port.PasswordVerifier,
	refreshTokens port.RefreshTokensRepo,
	signer port.TokenSigner,
	clock port.Clock,
	refreshTTL time.Duration,
) *LoginUseCase {
	return &LoginUseCase{
		users:         users,
		passwords:     passwords,
		refreshTokens: refreshTokens,
		signer:        signer,
		clock:         clock,
		refreshTTL:    refreshTTL,
	}
}

func (uc *LoginUseCase) Execute(
	ctx context.Context,
	cmd LoginCommand,
) (*LoginResult, error) {
	user, err := uc.users.FindByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if !uc.passwords.Compare(user.PasswordHash, cmd.Password) {
		return nil, domain.ErrInvalidCredentials
	}

	now := uc.clock.Now()

	userID, err := valueobject.ParseUserID(user.ID)
	if err != nil {
		return nil, err
	}

	refresh := entity.NewRefreshToken(
		valueobject.NewTokenID(),
		userID,
		now.Add(uc.refreshTTL),
	)

	if err := uc.refreshTokens.Save(ctx, refresh); err != nil {
		return nil, err
	}

	accessJWT, err := uc.signer.SignAccess(userID, now)
	if err != nil {
		return nil, err
	}

	refreshJWT, err := uc.signer.SignRefresh(refresh.ID, userID, now)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		AccessToken:  accessJWT,
		RefreshToken: refreshJWT,
	}, nil
}
