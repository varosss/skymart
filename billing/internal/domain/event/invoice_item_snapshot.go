package event

import "clirzy/billing/internal/domain/valueobject"

type InvoiceItemSnapshot struct {
	productID valueobject.ProductID
	amount    int64
	currency  string
	qty       int64
}

func NewInvoiceItemSnapshot(
	productID valueobject.ProductID,
	amount int64,
	currency string,
	qty int64,
) *InvoiceItemSnapshot {
	return &InvoiceItemSnapshot{
		productID: productID,
		amount:    amount,
		currency:  currency,
		qty:       qty,
	}
}

func (i *InvoiceItemSnapshot) ProductID() valueobject.ProductID {
	return i.productID
}

func (i *InvoiceItemSnapshot) Amount() int64 {
	return i.amount
}

func (i *InvoiceItemSnapshot) Currency() string {
	return i.currency
}

func (i *InvoiceItemSnapshot) Qty() int64 {
	return i.qty
}
