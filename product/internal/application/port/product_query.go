package port

import (
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type ProductDTO struct {
	ID       string
	SellerID string
	Price    int64
	Currency string
}

type ProductQuery interface {
	GetByID(ctx context.Context, productID valueobject.ProductID) (*ProductDTO, error)
	GetProducts(ctx context.Context, productIDs []valueobject.ProductID) ([]*ProductDTO, error)
}
