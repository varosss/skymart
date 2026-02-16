package valueobject

import "github.com/google/uuid"

type OrderID string

func NewOrderID() OrderID {
	return OrderID(uuid.New().String())
}

func ParseOrderID(id string) (OrderID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return OrderID(parsedUUID.String()), nil
}

func (id OrderID) String() string {
	return string(id)
}
