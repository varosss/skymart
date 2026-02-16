package event

import (
	"clirzy/payment/internal/domain/valueobject"
	"time"
)

type PaymentCreated struct {
	eventID    valueobject.EventID
	paymentID  valueobject.PaymentID
	invoiceID  valueobject.InvoiceID
	amount     int64
	currency   string
	occurredAt time.Time
}

func NewPaymentCreated(
	paymentID valueobject.PaymentID,
	invoiceID valueobject.InvoiceID,
	money valueobject.Money,
) *PaymentCreated {
	return &PaymentCreated{
		eventID:    valueobject.NewEventID(),
		paymentID:  paymentID,
		invoiceID:  invoiceID,
		amount:     money.Amount(),
		currency:   money.Currency(),
		occurredAt: time.Now(),
	}
}

func PaymentCreatedFromPrimitives(
	eventID valueobject.EventID,
	paymentID valueobject.PaymentID,
	invoiceID valueobject.InvoiceID,
	money valueobject.Money,
	occurredAt time.Time,
) *PaymentCreated {
	return &PaymentCreated{
		eventID:    eventID,
		paymentID:  paymentID,
		invoiceID:  invoiceID,
		amount:     money.Amount(),
		currency:   money.Currency(),
		occurredAt: occurredAt,
	}
}

func (e *PaymentCreated) ID() string {
	return e.eventID.String()
}

func (e *PaymentCreated) Type() string {
	return "payment.created"
}

func (e *PaymentCreated) AggregateID() string {
	return e.paymentID.String()
}

func (e *PaymentCreated) AggregateType() string {
	return "payment"
}

func (e *PaymentCreated) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *PaymentCreated) PaymentID() valueobject.PaymentID {
	return e.paymentID
}

func (e *PaymentCreated) InvoiceID() valueobject.InvoiceID {
	return e.invoiceID
}

func (e *PaymentCreated) Amount() int64 {
	return e.amount
}

func (e *PaymentCreated) Currency() string {
	return e.currency
}
