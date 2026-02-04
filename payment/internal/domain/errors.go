package domain

import "errors"

var (
	ErrPaymentNotFound              = errors.New("payment not found")
	ErrInvalidPaymentID             = errors.New("invalid payment id")
	ErrInvalidInvoiceID             = errors.New("invalid invoice id")
	ErrPaymentAlreadyExists         = errors.New("payment already exists")
	ErrInvalidPaymentState          = errors.New("invalid payment state")
	ErrCannotCancelSucceededPayment = errors.New("cannot cancel succeeded payment")
)
