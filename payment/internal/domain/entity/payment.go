package entity

import "clirzy/payment/internal/domain/valueobject"

type Payment struct {
	id         valueobject.PaymentID
	customerID valueobject.CustomerID
}

func NewPayment(customerID valueobject.CustomerID) *Payment {
	return &Payment{
		id:         valueobject.NewPaymentID(),
		customerID: customerID,
	}
}

func (p *Payment) ID() valueobject.PaymentID {
	return p.id
}

func (p *Payment) CustomerID() valueobject.CustomerID {
	return p.customerID
}
