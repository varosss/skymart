package event

import (
	"clirzy/billing/internal/domain/valueobject"
	"time"
)

type InvoiceCreated struct {
	eventID    valueobject.EventID
	invoiceID  valueobject.InvoiceID
	buyerID    valueobject.BuyerID
	total      int64
	currency   string
	items      []*InvoiceItemSnapshot
	occurredAt time.Time
}

func NewInvoiceCreated(
	invoiceID valueobject.InvoiceID,
	buyerID valueobject.BuyerID,
	total valueobject.Money,
	items []*InvoiceItemSnapshot,
) *InvoiceCreated {
	return &InvoiceCreated{
		eventID:    valueobject.NewEventID(),
		invoiceID:  invoiceID,
		buyerID:    buyerID,
		total:      total.Amount(),
		currency:   total.Currency(),
		items:      items,
		occurredAt: time.Now(),
	}
}

func InvoiceCreatedFromPrimitives(
	eventID valueobject.EventID,
	invoiceID valueobject.InvoiceID,
	buyerID valueobject.BuyerID,
	total valueobject.Money,
	items []*InvoiceItemSnapshot,
	occurredAt time.Time,
) *InvoiceCreated {
	return &InvoiceCreated{
		eventID:    eventID,
		invoiceID:  invoiceID,
		buyerID:    buyerID,
		total:      total.Amount(),
		currency:   total.Currency(),
		items:      items,
		occurredAt: occurredAt,
	}
}

func (e *InvoiceCreated) ID() string {
	return e.eventID.String()
}

func (*InvoiceCreated) Type() string {
	return "invoice.created"
}

func (e *InvoiceCreated) AggregateID() string {
	return e.invoiceID.String()
}

func (*InvoiceCreated) AggregateType() string {
	return "invoice"
}

func (e *InvoiceCreated) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *InvoiceCreated) InvoiceID() valueobject.InvoiceID {
	return e.invoiceID
}

func (e *InvoiceCreated) BuyerID() valueobject.BuyerID {
	return e.buyerID
}

func (e *InvoiceCreated) Total() int64 {
	return e.total
}

func (e *InvoiceCreated) Currency() string {
	return e.currency
}

func (e *InvoiceCreated) Items() []*InvoiceItemSnapshot {
	return e.items
}
