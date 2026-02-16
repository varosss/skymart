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
	Roles  []valueobject.RoleCode
}

type ValidateAccessTokenUseCase struct {
	verifier port.TokenVerifier
}

func NewValidateAccessTokenUseCase(verifier port.TokenVerifier) *ValidateAccessTokenUseCase {
	return &ValidateAccessTokenUseCase{
		verifier: verifier,
	}
}

func (uc *ValidateAccessTokenUseCase) Execute(
	ctx context.Context,
	cmd ValidateAccessTokenCommand,
) (*ValidateAccessTokenResult, error) {

	claims, err := uc.verifier.VerifyAccess(cmd.AccessToken)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	return &ValidateAccessTokenResult{
		UserID: claims.UserID,
		Roles:  claims.Roles,
	}, nil
}
