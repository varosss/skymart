package port

import (
	"clirzy/billing/internal/domain/valueobject"
	"context"
)

type InvoiceDTO struct {
	ID       string
	Amount   int64
	Currency string
	Status   string
}

type InvoiceQuery interface {
	GetByID(ctx context.Context, invoiceID valueobject.InvoiceID) (*InvoiceDTO, error)
}
