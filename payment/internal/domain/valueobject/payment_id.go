package valueobject

import "github.com/google/uuid"

type PaymentID string

func NewPaymentID() PaymentID {
	return PaymentID(uuid.NewString())
}

func ParsePaymentID(id string) (PaymentID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return PaymentID(parsedUUID.String()), nil
}
