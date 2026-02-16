package port

import (
	"clirzy/product/internal/domain/entity"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type ProductRepo interface {
	Save(ctx context.Context, product *entity.Product) error
	GetByID(ctx context.Context, id valueobject.ProductID) (*entity.Product, error)
}
