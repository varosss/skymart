package valueobject

import "github.com/google/uuid"

type InvoiceID string

func NewInvoiceID() InvoiceID {
	return InvoiceID(uuid.New().String())
}

func ParseInvoiceID(id string) (InvoiceID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return InvoiceID(parsedUUID.String()), nil
}

func (id InvoiceID) String() string {
	return string(id)
}
