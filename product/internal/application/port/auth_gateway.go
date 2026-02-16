package port

import (
	"context"
)

type AuthGateway interface {
	ValidateAccessToken(ctx context.Context, accessToken string) (userID string, roles []string, err error)
}
