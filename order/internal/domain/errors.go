package domain

import "errors"

var (
	ErrEmptyOrder          = errors.New("empty order")
	ErrInvalidProductID    = errors.New("invalid product id")
	ErrInvalidBuyerID      = errors.New("invalid buyer id")
	ErrInvalidSellerID     = errors.New("invalid seller id")
	ErrInvalidQuantity     = errors.New("invalid product quantity")
	ErrProductNotAvailable = errors.New("product is not available")
	ErrProductNotFound     = errors.New("product not found")
	ErrBuyerNotFound       = errors.New("buyer not found")
	ErrInactiveBuyer       = errors.New("buyer is inactive")
	ErrMixedCurrencies     = errors.New("currencies are mixed")
)
