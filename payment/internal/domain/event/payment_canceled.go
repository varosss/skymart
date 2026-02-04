package event

import (
	"clirzy/payment/internal/domain/valueobject"
	"time"
)

type PaymentCanceled struct {
	eventID    valueobject.EventID
	paymentID  valueobject.PaymentID
	invoiceID  valueobject.InvoiceID
	occurredAt time.Time
}

func NewPaymentCanceled(
	paymentID valueobject.PaymentID,
	invoiceID valueobject.InvoiceID,
) PaymentCanceled {
	return PaymentCanceled{
		eventID:    valueobject.NewEventID(),
		paymentID:  paymentID,
		invoiceID:  invoiceID,
		occurredAt: time.Now(),
	}
}

func (e PaymentCanceled) ID() string {
	return e.eventID.String()
}

func (e PaymentCanceled) Type() string {
	return "payment.cancelled"
}

func (e PaymentCanceled) AggregateID() string {
	return e.eventID.String()
}

func (e PaymentCanceled) AggregateType() string {
	return "payment"
}

func (e PaymentCanceled) OccurredAt() time.Time {
	return e.occurredAt
}

func (e PaymentCanceled) PaymentID() valueobject.PaymentID {
	return e.paymentID
}

func (e PaymentCanceled) InvoiceID() valueobject.InvoiceID {
	return e.invoiceID
}
