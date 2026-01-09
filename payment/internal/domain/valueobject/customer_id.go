package valueobject

import "github.com/google/uuid"

type CustomerID string

func ParseCustomerID(id string) (CustomerID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return CustomerID(parsedUUID.String()), nil
}
