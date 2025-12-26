package valueobject

import "github.com/google/uuid"

type BuyerID string

func NewBuyerID() BuyerID {
	return BuyerID(uuid.NewString())
}

func ToBuyerID(id string) (BuyerID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return BuyerID(parsedUUID.String()), nil
}
