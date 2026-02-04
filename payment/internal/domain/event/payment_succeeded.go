package event

import (
	"clirzy/payment/internal/domain/valueobject"
	"time"
)

type PaymentSucceeded struct {
	eventID    valueobject.EventID
	paymentID  valueobject.PaymentID
	invoiceID  valueobject.InvoiceID
	amount     int64
	currency   string
	occurredAt time.Time
}

func NewPaymentSucceeded(
	paymentID valueobject.PaymentID,
	invoiceID valueobject.InvoiceID,
	money valueobject.Money,
) PaymentSucceeded {
	return PaymentSucceeded{
		eventID:    valueobject.NewEventID(),
		paymentID:  paymentID,
		invoiceID:  invoiceID,
		amount:     money.Amount(),
		currency:   money.Currency(),
		occurredAt: time.Now(),
	}
}

func (e PaymentSucceeded) ID() string {
	return e.eventID.String()
}

func (e PaymentSucceeded) Type() string {
	return "payment.succeeded"
}

func (e PaymentSucceeded) AggregateID() string {
	return e.paymentID.String()
}

func (e PaymentSucceeded) AggregateType() string {
	return "payment"
}

func (e PaymentSucceeded) OccurredAt() time.Time {
	return e.occurredAt
}

func (e PaymentSucceeded) PaymentID() valueobject.PaymentID {
	return e.paymentID
}

func (e PaymentSucceeded) InvoiceID() valueobject.InvoiceID {
	return e.invoiceID
}

func (e PaymentSucceeded) Amount() int64 {
	return e.amount
}

func (e PaymentSucceeded) Currency() string {
	return e.currency
}
