package valueobject

import "github.com/google/uuid"

type SellerID string

func NewSellerID() SellerID {
	return SellerID(uuid.NewString())
}

func ToSellerID(id string) (SellerID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return SellerID(parsedUUID.String()), nil
}
