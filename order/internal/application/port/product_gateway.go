package port

import (
	"clirzy/order/internal/domain/valueobject"
	"context"
)

type ProductDTO struct {
	ID          valueobject.ProductID
	SellerID    valueobject.SellerID
	Price       valueobject.Money
	IsPublished bool
}

type ProductGateway interface {
	GetProducts(ctx context.Context, productIDs []valueobject.ProductID) ([]*ProductDTO, error)
}
