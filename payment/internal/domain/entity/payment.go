package entity

import (
	"clirzy/payment/internal/domain"
	"clirzy/payment/internal/domain/event"
	"clirzy/payment/internal/domain/valueobject"
)

type Payment struct {
	id        valueobject.PaymentID
	invoiceID valueobject.InvoiceID
	amount    valueobject.Money
	status    valueobject.PaymentStatus

	events []event.Event
}

func NewPayment(
	invoiceID valueobject.InvoiceID,
	amount valueobject.Money,
) (*Payment, error) {

	p := &Payment{
		id:        valueobject.NewPaymentID(),
		invoiceID: invoiceID,
		amount:    amount,
		status:    valueobject.PaymentStatusPending,
	}

	p.addEvent(event.NewPaymentCreated(
		p.id,
		p.invoiceID,
		p.amount,
	))

	return p, nil
}

func (p *Payment) ID() valueobject.PaymentID {
	return p.id
}

func (p *Payment) MarkSucceeded() error {
	if p.status != valueobject.PaymentStatusPending {
		return domain.ErrInvalidPaymentState
	}

	p.status = valueobject.PaymentStatusSucceeded

	p.addEvent(event.NewPaymentSucceeded(
		p.id,
		p.invoiceID,
		p.amount,
	))

	return nil
}

func (p *Payment) Cancel() error {
	if p.status == valueobject.PaymentStatusSucceeded {
		return domain.ErrCannotCancelSucceededPayment
	}

	if p.status == valueobject.PaymentStatusCanceled {
		return nil
	}

	p.status = valueobject.PaymentStatusCanceled

	p.addEvent(event.NewPaymentCanceled(
		p.id,
		p.invoiceID,
	))

	return nil
}

func (p *Payment) PullEvents() []event.Event {
	events := p.events
	p.events = nil
	return events
}

func (p *Payment) addEvent(e event.Event) {
	p.events = append(p.events, e)
}
