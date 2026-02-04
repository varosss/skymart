package domain

import "errors"

var (
	ErrEmptyInvoice            = errors.New("invoice is empty")
	ErrMixedCurrencies         = errors.New("currencies are mixed")
	ErrInvalidQuantity         = errors.New("invalid quantity")
	ErrInvalidOrderID          = errors.New("invalid order id")
	ErrInvalidBuyerID          = errors.New("invalid buyer id")
	ErrInvalidProductID        = errors.New("invalid product id")
	ErrInvalidInvoiceID        = errors.New("invalid invoice id")
	ErrInvoiceNotFound         = errors.New("invoice not found")
	ErrInvalidPaymentAmount    = errors.New("invalid payment amount")
	ErrInvalidInvoiceState     = errors.New("invalid invoice state")
	ErrCannotCancelPaidInvoice = errors.New("cannot cancel paid invoice")
)
