package port

import "clirzy/auth/internal/domain/valueobject"

type AccessTokenClaims struct {
	UserID valueobject.UserID
	Roles  []valueobject.RoleCode
}

type TokenVerifier interface {
	VerifyAccess(token string) (*AccessTokenClaims, error)
	VerifyRefresh(token string) (valueobject.TokenID, valueobject.UserID, error)
}
