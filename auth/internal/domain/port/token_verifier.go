package port

import "clirzy/auth/internal/domain/valueobject"

type TokenVerifier interface {
	VerifyAccess(token string) (valueobject.UserID, error)
	VerifyRefresh(token string) (valueobject.TokenID, valueobject.UserID, error)
}
