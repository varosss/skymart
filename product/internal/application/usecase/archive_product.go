package usecase

import (
	aport "clirzy/product/internal/application/port"
	"clirzy/product/internal/application/service"
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/port"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type ArchiveProductCommand struct {
	SellerID  string
	ProductID string
}

type ArchiveProductUseCase struct {
	sellers  aport.SellerQuery
	products port.ProductRepo
	access   *service.ProductAccessService
	bus      aport.EventBus
}

func NewArchiveProductUseCase(
	sellers aport.SellerQuery,
	products port.ProductRepo,
	access *service.ProductAccessService,
	bus aport.EventBus,
) *ArchiveProductUseCase {
	return &ArchiveProductUseCase{
		sellers:  sellers,
		products: products,
		access:   access,
		bus:      bus,
	}
}

func (uc *ArchiveProductUseCase) Execute(
	ctx context.Context,
	cmd ArchiveProductCommand,
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

	if err := product.Archive(); err != nil {
		return err
	}

	if err := uc.products.Save(ctx, product); err != nil {
		return err
	}

	return uc.bus.Publish(ctx, product.PullEvents())
}
