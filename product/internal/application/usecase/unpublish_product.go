package usecase

import (
	aport "clirzy/product/internal/application/port"
	"clirzy/product/internal/application/service"
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/port"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type UnpublishProductCommand struct {
	SellerID  string
	ProductID string
}

type UnpublishProductUseCase struct {
	sellers  aport.SellerService
	products port.ProductsRepo
	access   *service.ProductAccessService
}

func NewUnpublishProductUseCase(
	sellers aport.SellerService,
	products port.ProductsRepo,
	access *service.ProductAccessService,
) *UnpublishProductUseCase {
	return &UnpublishProductUseCase{
		sellers:  sellers,
		products: products,
		access:   access,
	}
}

func (uc *UnpublishProductUseCase) Execute(
	ctx context.Context,
	cmd UnpublishProductCommand,
) error {
	sellerID, err := valueobject.ToSellerID(cmd.SellerID)
	if err != nil {
		return domain.ErrInvalidSellerID
	}

	productID, err := valueobject.ToProductID(cmd.ProductID)
	if err != nil {
		return domain.ErrInvalidProductID
	}

	product, err := uc.access.LoadForSeller(ctx, sellerID, productID)
	if err != nil {
		return err
	}

	if err := product.Unpublish(); err != nil {
		return err
	}

	return uc.products.Save(ctx, product)
}
