package usecase

import (
	"clirzy/auth/internal/domain"
	"clirzy/auth/internal/domain/entity"
	"clirzy/auth/internal/domain/port"
	"clirzy/auth/internal/domain/valueobject"
	"context"
	"time"
)

type RefreshTokenCommand struct {
	RefreshToken string
}

type RefreshTokenResult struct {
	AccessToken  string
	RefreshToken string
}

type RefreshTokenUseCase struct {
	refreshTokens port.RefreshTokensRepo
	verifier      port.TokenVerifier
	signer        port.TokenSigner
	clock         port.Clock
	refreshTTL    time.Duration
}

func NewRefreshTokenUseCase(
	refreshTokens port.RefreshTokensRepo,
	verifier port.TokenVerifier,
	signer port.TokenSigner,
	clock port.Clock,
	refreshTTL time.Duration,
) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		refreshTokens: refreshTokens,
		verifier:      verifier,
		signer:        signer,
		clock:         clock,
		refreshTTL:    refreshTTL,
	}
}

func (uc *RefreshTokenUseCase) Execute(
	ctx context.Context,
	cmd RefreshTokenCommand,
) (*RefreshTokenResult, error) {

	tokenID, userID, err := uc.verifier.VerifyRefresh(cmd.RefreshToken)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	stored, err := uc.refreshTokens.Get(ctx, tokenID)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	now := uc.clock.Now()

	if stored.Revoked || stored.IsExpired(now) {
		return nil, domain.ErrInvalidToken
	}

	stored.Revoke()
	if err := uc.refreshTokens.Revoke(ctx, stored.ID); err != nil {
		return nil, err
	}

	newRefresh := entity.NewRefreshToken(
		valueobject.NewTokenID(),
		userID,
		now.Add(uc.refreshTTL),
	)

	if err := uc.refreshTokens.Save(ctx, newRefresh); err != nil {
		return nil, err
	}

	accessJWT, err := uc.signer.SignAccess(userID, now)
	if err != nil {
		return nil, err
	}

	refreshJWT, err := uc.signer.SignRefresh(newRefresh.ID, userID, now)
	if err != nil {
		return nil, err
	}

	return &RefreshTokenResult{
		AccessToken:  accessJWT,
		RefreshToken: refreshJWT,
	}, nil
}
