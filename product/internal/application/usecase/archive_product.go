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
	sellers   aport.SellerQuery
	products  port.ProductsRepo
	access    *service.ProductAccessService
	publisher aport.EventPublisher
}

func NewArchiveProductUseCase(
	sellers aport.SellerQuery,
	products port.ProductsRepo,
	access *service.ProductAccessService,
	publisher aport.EventPublisher,
) *ArchiveProductUseCase {
	return &ArchiveProductUseCase{
		sellers:   sellers,
		products:  products,
		access:    access,
		publisher: publisher,
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

	return uc.publisher.Publish(product.PullEvents())
}
