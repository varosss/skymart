package port

import (
	"clirzy/order/internal/domain/valueobject"
	"context"
)

type BuyerDTO struct {
	ID       string
	IsActive bool
}

type BuyerGateway interface {
	GetByID(ctx context.Context, buyerID valueobject.BuyerID) (*BuyerDTO, error)
}
