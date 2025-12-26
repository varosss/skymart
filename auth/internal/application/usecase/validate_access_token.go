package usecase

import (
	"clirzy/auth/internal/domain"
	"clirzy/auth/internal/domain/port"
	"clirzy/auth/internal/domain/valueobject"
	"context"
)

type ValidateAccessTokenCommand struct {
	AccessToken string
}

type ValidateAccessTokenResult struct {
	UserID valueobject.UserID
}

type ValidateAccessTokenUseCase struct {
	verifier port.TokenVerifier
}

func (uc *ValidateAccessTokenUseCase) Execute(
	ctx context.Context,
	cmd ValidateAccessTokenCommand,
) (*ValidateAccessTokenResult, error) {

	userID, err := uc.verifier.VerifyAccess(cmd.AccessToken)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	return &ValidateAccessTokenResult{
		UserID: userID,
	}, nil
}
