package usecase

import (
	aport "clirzy/billing/internal/application/port"
	"clirzy/billing/internal/domain"
	"clirzy/billing/internal/domain/port"
	"clirzy/billing/internal/domain/valueobject"
	"context"
)

type PayInvoiceCommand struct {
	InvoiceID string
	Amount    int64
	Currency  string
}

type PayInvoiceUseCase struct {
	invoices port.InvoiceRepo
	bus      aport.EventBus
}

func NewPayInvoiceUseCase(
	invoices port.InvoiceRepo,
) *PayInvoiceUseCase {
	return &PayInvoiceUseCase{
		invoices: invoices,
	}
}

func (uc *PayInvoiceUseCase) Execute(
	ctx context.Context,
	cmd PayInvoiceCommand,
) error {

	invoiceID, err := valueobject.ParseInvoiceID(cmd.InvoiceID)
	if err != nil {
		return domain.ErrInvalidInvoiceID
	}

	invoice, err := uc.invoices.FindByID(ctx, invoiceID)
	if err != nil {
		return err
	}
	if invoice == nil {
		return domain.ErrInvoiceNotFound
	}

	if invoice.Total().Amount() != cmd.Amount ||
		invoice.Total().Currency() != cmd.Currency {
		return domain.ErrInvalidPaymentAmount
	}

	if err := invoice.MarkPaid(); err != nil {
		return err
	}

	if err := uc.invoices.Save(ctx, invoice); err != nil {
		return err
	}

	if err := uc.bus.Publish(ctx, invoice.PullEvents()); err != nil {
		return err
	}

	return nil
}
