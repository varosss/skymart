package event

import (
	"clirzy/billing/internal/domain/valueobject"
	"time"
)

type InvoicePaid struct {
	eventID    valueobject.EventID
	invoiceID  valueobject.InvoiceID
	buyerID    valueobject.BuyerID
	amount     int64
	currency   string
	occurredAt time.Time
}

func NewInvoicePaid(
	id valueobject.InvoiceID,
	buyerID valueobject.BuyerID,
	total valueobject.Money,
) *InvoicePaid {
	return &InvoicePaid{
		eventID:    valueobject.NewEventID(),
		invoiceID:  id,
		buyerID:    buyerID,
		amount:     total.Amount(),
		currency:   total.Currency(),
		occurredAt: time.Now(),
	}
}

func (e *InvoicePaid) ID() string {
	return e.eventID.String()
}

func (*InvoicePaid) Type() string {
	return "invoice.paid"
}

func (e *InvoicePaid) AggregateID() string {
	return e.invoiceID.String()
}

func (e *InvoicePaid) AggregateType() string {
	return "invoice"
}

func (e *InvoicePaid) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *InvoicePaid) InvoiceID() valueobject.InvoiceID {
	return e.invoiceID
}

func (e *InvoicePaid) BuyerID() valueobject.BuyerID {
	return e.buyerID
}

func (e *InvoicePaid) Amount() int64 {
	return e.amount
}

func (e *InvoicePaid) Currency() string {
	return e.currency
}
