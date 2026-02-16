package valueobject

import "github.com/google/uuid"

type SellerID string

func ParseSellerID(id string) (SellerID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return SellerID(parsedUUID.String()), nil
}

func (id SellerID) String() string {
	return string(id)
}
