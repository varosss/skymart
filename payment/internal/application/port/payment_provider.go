package port

import (
	"context"
)

type CreatePaymentRequest struct {
	PaymentID         string // inner ID
	Amount            int64  // в минимальных единицах (копейки, центы)
	Currency          string // ISO 4217
	Description       string
	CustomerID        string
	CustomerEmail     string
	CustomerPhone     string
	ReturnURL         string // success URL
	CancelURL         string
	PaymentMethod     string
	SavePaymentMethod bool
	Metadata          map[string]string
}

type CreatePaymentResponse struct {
	Action       string
	ProviderData map[string]string
}

type PaymentProvider interface {
	CreatePayment(
		ctx context.Context,
		req CreatePaymentRequest,
	) (*CreatePaymentResponse, error)

	CancelPayment(
		ctx context.Context,
		paymentID string,
	) error
}
