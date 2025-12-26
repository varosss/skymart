package valueobject

import "github.com/google/uuid"

type TokenID string

func NewTokenID() TokenID {
	return TokenID(uuid.NewString())
}
