package port

import (
	"clirzy/payment/internal/domain/valueobject"
	"context"
)

type InvoiceDTO struct {
	ID       string
	Amount   int64
	Currency string
	Status   string
}

type BillingGateway interface {
	GetInvoiceByID(ctx context.Context, invoiceID valueobject.InvoiceID) (*InvoiceDTO, error)
}
