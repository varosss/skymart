package auth

import "context"

type TokenValidator interface {
	ValidateAccessToken(ctx context.Context, accessToken string) (userID string, roles []string, err error)
}
