package usecase

import (
	"clirzy/auth/internal/domain"
	"clirzy/auth/internal/domain/entity"
	"clirzy/auth/internal/domain/port"
	"clirzy/auth/internal/domain/valueobject"
	"context"
	"time"
)

type RefreshCommand struct {
	RefreshToken string
}

type RefreshResult struct {
	AccessToken  string
	RefreshToken string
}

type RefreshTokensUseCase struct {
	refreshRepo port.RefreshTokensRepo
	verifier    port.TokenVerifier
	signer      port.TokenSigner
	clock       port.Clock
	refreshTTL  time.Duration
}

func (uc *RefreshTokensUseCase) Execute(
	ctx context.Context,
	cmd RefreshCommand,
) (*RefreshResult, error) {

	tokenID, userID, err := uc.verifier.VerifyRefresh(cmd.RefreshToken)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	stored, err := uc.refreshRepo.Get(ctx, tokenID)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	now := uc.clock.Now()

	if stored.Revoked || stored.IsExpired(now) {
		return nil, domain.ErrInvalidToken
	}

	stored.Revoke()
	if err := uc.refreshRepo.Revoke(ctx, stored.ID); err != nil {
		return nil, err
	}

	newRefresh := entity.NewRefreshToken(
		valueobject.NewTokenID(),
		userID,
		now.Add(uc.refreshTTL),
	)

	if err := uc.refreshRepo.Save(ctx, newRefresh); err != nil {
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

	return &RefreshResult{
		AccessToken:  accessJWT,
		RefreshToken: refreshJWT,
	}, nil
}
