package usecase

import (
	"clirzy/auth/internal/domain"
	"clirzy/auth/internal/domain/port"
	"context"
)

type LogoutCommand struct {
	RefreshToken string
}

type LogoutUseCase struct {
	verifier    port.TokenVerifier
	refreshRepo port.RefreshTokensRepo
}

func (uc *LogoutUseCase) Execute(
	ctx context.Context,
	cmd LogoutCommand,
) error {

	tokenID, _, err := uc.verifier.VerifyRefresh(cmd.RefreshToken)
	if err != nil {
		return domain.ErrInvalidToken
	}

	return uc.refreshRepo.Revoke(ctx, tokenID)
}
