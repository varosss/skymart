package valueobject

import "github.com/google/uuid"

type TokenID string

func NewTokenID() TokenID {
	return TokenID(uuid.NewString())
}

func (id TokenID) String() string {
	return string(id)
}
