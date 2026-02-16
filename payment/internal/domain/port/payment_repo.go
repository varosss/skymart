package port

import (
	"clirzy/payment/internal/domain/entity"
	"clirzy/payment/internal/domain/valueobject"
	"context"
)

type PaymentRepo interface {
	Save(ctx context.Context, payment *entity.Payment) error
	FindByID(ctx context.Context, paymentID valueobject.PaymentID) (*entity.Payment, error)
	FindByInvoiceID(ctx context.Context, invoiceID valueobject.InvoiceID) (*entity.Payment, error)
}
