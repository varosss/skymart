package fakes

import (
	"clirzy/auth/internal/domain/valueobject"
	"time"
)

type FakeTokenSigner struct {
	AccessErr  error
	RefreshErr error
}

func (f *FakeTokenSigner) SignAccess(
	userID valueobject.UserID,
	roles []valueobject.RoleCode,
	now time.Time,
) (string, error) {
	if f.AccessErr != nil {
		return "", f.AccessErr
	}
	return "access.jwt", nil
}

func (f *FakeTokenSigner) SignRefresh(
	tokenID valueobject.TokenID,
	userID valueobject.UserID,
	now time.Time,
) (string, error) {
	if f.RefreshErr != nil {
		return "", f.RefreshErr
	}
	return "refresh.jwt", nil
}
