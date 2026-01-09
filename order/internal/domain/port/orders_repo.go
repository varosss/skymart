package port

import (
	"clirzy/order/internal/domain/entity"
	"clirzy/order/internal/domain/valueobject"
	"context"
)

type OrdersRepo interface {
	Save(ctx context.Context, o *entity.Order) error
	GetByID(ctx context.Context, id valueobject.OrderID) (*entity.Order, error)
}
