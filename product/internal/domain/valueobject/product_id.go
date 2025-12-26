package valueobject

import "github.com/google/uuid"

type ProductID string

func NewProductID() ProductID {
	return ProductID(uuid.New().String())
}

func ToProductID(id string) (ProductID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return ProductID(parsedUUID.String()), nil
}
