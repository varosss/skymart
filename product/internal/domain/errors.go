package domain

import "errors"

var (
	ErrInvalidSellerID                = errors.New("invalid seller id")
	ErrInvalidProductID               = errors.New("invalid product id")
	ErrSellerNotFound                 = errors.New("seller not found")
	ErrSellerInactive                 = errors.New("seller is inactive")
	ErrProductNotFound                = errors.New("product not found")
	ErrInvalidProductStatusTransition = errors.New("invalid product status transition")
	ErrProductNotOwnedBySeller        = errors.New("product is not owned by seller")
	ErrCannotChangeArchivedProduct    = errors.New("cannot change archived product")
	ErrInvalidTitle                   = errors.New("invalid title")
	ErrInvalidPrice                   = errors.New("invalid price")
)
