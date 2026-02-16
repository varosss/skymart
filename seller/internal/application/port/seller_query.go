package port

import (
	"clirzy/seller/internal/domain/valueobject"
	"context"
)

type SellerDTO struct {
	ID       string
	UserID   string
	IsActive bool
}

type SellerQuery interface {
	GetByID(ctx context.Context, sellerID valueobject.SellerID) (*SellerDTO, error)
	GetByUserID(ctx context.Context, userID valueobject.UserID) (*SellerDTO, error)
}
