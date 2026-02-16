package port

import (
	"clirzy/billing/internal/domain/entity"
	"clirzy/billing/internal/domain/valueobject"
	"context"
)

type InvoiceRepo interface {
	Save(ctx context.Context, invoice *entity.Invoice) error
	FindByID(ctx context.Context, invoiceID valueobject.InvoiceID) (*entity.Invoice, error)
}
