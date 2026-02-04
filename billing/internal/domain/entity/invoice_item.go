package entity

import (
	"clirzy/billing/internal/domain"
	"clirzy/billing/internal/domain/valueobject"
)

type InvoiceItem struct {
	productID valueobject.ProductID
	price     valueobject.Money
	qty       int64
}

func NewInvoiceItem(
	productID valueobject.ProductID,
	price valueobject.Money,
	qty int64,
) (InvoiceItem, error) {

	if qty <= 0 {
		return InvoiceItem{}, domain.ErrInvalidQuantity
	}

	return InvoiceItem{
		productID: productID,
		price:     price,
		qty:       qty,
	}, nil
}
