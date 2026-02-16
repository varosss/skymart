package usecase

import (
	aport "clirzy/billing/internal/application/port"
	"clirzy/billing/internal/domain"
	"clirzy/billing/internal/domain/entity"
	"clirzy/billing/internal/domain/port"
	"clirzy/billing/internal/domain/valueobject"
	"context"
)

type CreateInvoiceItem struct {
	ProductID valueobject.ProductID
	Money     valueobject.Money
	Qty       int64
}

type CreateInvoiceCommand struct {
	BuyerID valueobject.BuyerID
	OrderID valueobject.OrderID
	Items   []CreateInvoiceItem
}

type CreateInvoiceUseCase struct {
	invoices port.InvoiceRepo
	bus      aport.EventBus
}

func NewCreateInvoiceUseCase(
	invoices port.InvoiceRepo,
) *CreateInvoiceUseCase {
	return &CreateInvoiceUseCase{
		invoices: invoices,
	}
}

func (uc *CreateInvoiceUseCase) Execute(
	ctx context.Context,
	cmd CreateInvoiceCommand,
) (valueobject.InvoiceID, error) {
	if len(cmd.Items) == 0 {
		return "", domain.ErrEmptyInvoice
	}

	items := make([]entity.InvoiceItem, 0, len(cmd.Items))

	for _, i := range cmd.Items {
		item, err := entity.NewInvoiceItem(
			i.ProductID,
			i.Money,
			i.Qty,
		)
		if err != nil {
			return "", err
		}

		items = append(items, item)
	}

	invoice, err := entity.NewInvoice(
		cmd.BuyerID,
		cmd.OrderID,
		items,
	)
	if err != nil {
		return "", err
	}

	if err := uc.invoices.Save(ctx, invoice); err != nil {
		return "", err
	}

	if err := uc.bus.Publish(ctx, invoice.PullEvents()); err != nil {
		return "", err
	}

	return invoice.ID(), nil
}
