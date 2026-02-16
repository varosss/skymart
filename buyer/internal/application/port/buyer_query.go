package port

import (
	"clirzy/buyer/internal/domain/valueobject"
	"context"
)

type BuyerDTO struct {
	ID       string
	IsActive bool
}

type BuyerQuery interface {
	GetByID(ctx context.Context, buyerID valueobject.BuyerID) (*BuyerDTO, error)
}
