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
	UserID      valueobject.UserID
	Title       string
	Description string
	Money       valueobject.Money
}

type CreateProductUseCase struct {
	sellers  aport.SellerGateway
	products port.ProductRepo
	bus      aport.EventBus
}

func NewCreateProductUseCase(
	sellers aport.SellerGateway,
	products port.ProductRepo,
	bus aport.EventBus,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		sellers:  sellers,
		products: products,
		bus:      bus,
	}
}

func (uc *CreateProductUseCase) Execute(
	ctx context.Context,
	cmd CreateProductCommand,
) (valueobject.ProductID, error) {
	seller, err := uc.sellers.GetByUserID(ctx, cmd.UserID)
	if err != nil {
		return "", err
	}
	if !seller.IsActive {
		return "", domain.ErrSellerInactive
	}

	product := entity.NewProduct(
		valueobject.NewProductID(),
		seller.ID,
		cmd.Title,
		cmd.Description,
		cmd.Money,
		valueobject.StatusDraft,
	)

	if err := uc.products.Save(ctx, product); err != nil {
		return "", err
	}

	if err := uc.bus.Publish(ctx, product.PullEvents()); err != nil {
		return "", err
	}

	return product.ID(), nil
}
