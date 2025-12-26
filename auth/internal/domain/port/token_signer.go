package port

import (
	"clirzy/auth/internal/domain/valueobject"
	"time"
)

type TokenSigner interface {
	SignAccess(userID valueobject.UserID, now time.Time) (string, error)
	SignRefresh(tokenID valueobject.TokenID, userID valueobject.UserID, now time.Time) (string, error)
}
