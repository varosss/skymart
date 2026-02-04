package event

import (
	"clirzy/billing/internal/domain/valueobject"
	"time"
)

type InvoiceCanceled struct {
	eventID    valueobject.EventID
	invoiceID  valueobject.InvoiceID
	buyerID    valueobject.BuyerID
	occurredAt time.Time
}

func NewInvoiceCanceled(
	invoiceID valueobject.InvoiceID,
	buyerID valueobject.BuyerID,
) *InvoiceCanceled {
	return &InvoiceCanceled{
		eventID:    valueobject.NewEventID(),
		invoiceID:  invoiceID,
		buyerID:    buyerID,
		occurredAt: time.Now(),
	}
}

func (e *InvoiceCanceled) Type() string {
	return "invoice.canceled"
}

func (e *InvoiceCanceled) ID() string {
	return e.eventID.String()
}

func (e *InvoiceCanceled) AggregateID() string {
	return e.invoiceID.String()
}

func (e *InvoiceCanceled) AggregateType() string {
	return "invoice"
}

func (e *InvoiceCanceled) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *InvoiceCanceled) InvoiceID() valueobject.InvoiceID {
	return e.invoiceID
}

func (e *InvoiceCanceled) BuyerID() valueobject.BuyerID {
	return e.buyerID
}
