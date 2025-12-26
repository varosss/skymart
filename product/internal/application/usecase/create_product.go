package usecase

import (
	aport "clirzy/product/internal/application/port"
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/entity"
	"clirzy/product/internal/domain/port"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type CreateProductCommand struct {
	SellerID    string
	Title       string
	Description string
	Price       int64
	Currency    string
}

type CreateProductUseCase struct {
	sellers   aport.SellerService
	products  port.ProductsRepo
	publisher aport.EventPublisher
}

func NewCreateProductUseCase(
	sellers aport.SellerService,
	products port.ProductsRepo,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		sellers:  sellers,
		products: products,
	}
}

func (uc *CreateProductUseCase) Execute(
	ctx context.Context,
	cmd CreateProductCommand,
) (valueobject.ProductID, error) {

	sellerID, err := valueobject.ToSellerID(cmd.SellerID)
	if err != nil {
		return "", domain.ErrInvalidSellerID
	}

	exists, err := uc.sellers.Exists(ctx, sellerID)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", domain.ErrSellerNotFound
	}

	active, err := uc.sellers.IsActive(ctx, sellerID)
	if err != nil {
		return "", err
	}
	if !active {
		return "", domain.ErrSellerInactive
	}

	product := entity.NewProduct(
		valueobject.NewProductID(),
		sellerID,
		cmd.Title,
		cmd.Description,
		valueobject.Money{Amount: cmd.Price, Currency: cmd.Currency},
		valueobject.StatusDraft,
	)

	if err := uc.products.Save(ctx, product); err != nil {
		return "", err
	}

	uc.publisher.Publish(product.PullEvents())

	return product.ID(), nil
}
