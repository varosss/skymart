package valueobject

import "github.com/google/uuid"

type BuyerID string

func ParseBuyerID(id string) (BuyerID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return BuyerID(parsedUUID.String()), nil
}

func (id BuyerID) String() string {
	return string(id)
}
