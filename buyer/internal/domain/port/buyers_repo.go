package port

import (
	"clirzy/buyer/internal/domain/entity"
	"clirzy/buyer/internal/domain/valueobject"
	"context"
)

type BuyersRepo interface {
	Save(ctx context.Context, buyer *entity.Buyer) error
	FindByID(ctx context.Context, buyerID valueobject.BuyerID) (*entity.Buyer, error)
	FindByUserID(ctx context.Context, userID valueobject.UserID) (*entity.Buyer, error)
}
