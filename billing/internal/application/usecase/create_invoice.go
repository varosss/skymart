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
	ProductID string
	Price     int64
	Currency  string
	Qty       int64
}

type CreateInvoiceCommand struct {
	BuyerID string
	OrderID string
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

	buyerID, err := valueobject.ParseBuyerID(cmd.BuyerID)
	if err != nil {
		return "", domain.ErrInvalidBuyerID
	}

	orderID, err := valueobject.ParseOrderID(cmd.OrderID)
	if err != nil {
		return "", domain.ErrInvalidOrderID
	}

	if len(cmd.Items) == 0 {
		return "", domain.ErrEmptyInvoice
	}

	items := make([]entity.InvoiceItem, 0, len(cmd.Items))

	for _, i := range cmd.Items {
		productID, err := valueobject.ParseProductID(i.ProductID)
		if err != nil {
			return "", domain.ErrInvalidProductID
		}

		money, err := valueobject.NewMoney(i.Price, i.Currency)
		if err != nil {
			return "", err
		}

		item, err := entity.NewInvoiceItem(
			productID,
			money,
			i.Qty,
		)
		if err != nil {
			return "", err
		}

		items = append(items, item)
	}

	invoice, err := entity.NewInvoice(
		buyerID,
		orderID,
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
