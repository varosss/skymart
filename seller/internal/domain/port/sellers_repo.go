package port

import (
	"clirzy/seller/internal/domain/entity"
	"clirzy/seller/internal/domain/valueobject"
	"context"
)

type SellersRepo interface {
	Save(ctx context.Context, buyer *entity.Seller) error
	FindByID(ctx context.Context, buyerID valueobject.SellerID) (*entity.Seller, error)
	FindByUserID(ctx context.Context, userID valueobject.UserID) (*entity.Seller, error)
}
