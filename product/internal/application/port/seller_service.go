package port

import (
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type SellerService interface {
	Exists(ctx context.Context, sellerID valueobject.SellerID) (bool, error)
	IsActive(ctx context.Context, sellerID valueobject.SellerID) (bool, error)
}
