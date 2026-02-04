package usecase

import (
	aport "clirzy/product/internal/application/port"
	"clirzy/product/internal/application/service"
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/port"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type PublishProductCommand struct {
	SellerID  string
	ProductID string
}

type PublishProductUseCase struct {
	sellers  aport.SellerQuery
	products port.ProductRepo
	access   *service.ProductAccessService
	bus      aport.EventBus
}

func NewPublishProductUseCase(
	sellers aport.SellerQuery,
	products port.ProductRepo,
	access *service.ProductAccessService,
	bus aport.EventBus,
) *PublishProductUseCase {
	return &PublishProductUseCase{
		sellers:  sellers,
		products: products,
		access:   access,
		bus:      bus,
	}
}

func (uc *PublishProductUseCase) Execute(
	ctx context.Context,
	cmd PublishProductCommand,
) error {
	sellerID, err := valueobject.ParseSellerID(cmd.SellerID)
	if err != nil {
		return domain.ErrInvalidSellerID
	}

	productID, err := valueobject.ParseProductID(cmd.ProductID)
	if err != nil {
		return domain.ErrInvalidProductID
	}

	product, err := uc.access.LoadForSeller(ctx, sellerID, productID)
	if err != nil {
		return err
	}

	if err := product.Publish(); err != nil {
		return err
	}

	if err := uc.products.Save(ctx, product); err != nil {
		return err
	}

	return uc.bus.Publish(ctx, product.PullEvents())
}
