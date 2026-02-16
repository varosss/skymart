package port

import (
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type SellerDTO struct {
	ID       valueobject.SellerID
	IsActive bool
}

type SellerGateway interface {
	GetByID(ctx context.Context, sellerID valueobject.SellerID) (*SellerDTO, error)
	GetByUserID(ctx context.Context, userID valueobject.UserID) (*SellerDTO, error)
}
