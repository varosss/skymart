package usecase

import (
	"clirzy/product/internal/application/service"
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/port"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type UpdateProductInfoCommand struct {
	SellerID    string
	ProductID   string
	Title       string
	Description string
}

type UpdateProductInfoUseCase struct {
	access   *service.ProductAccessService
	products port.ProductsRepo
}

func NewUpdateProductInfoUseCase(
	access *service.ProductAccessService,
	products port.ProductsRepo,
) *UpdateProductInfoUseCase {
	return &UpdateProductInfoUseCase{access, products}
}

func (uc *UpdateProductInfoUseCase) Execute(
	ctx context.Context,
	cmd UpdateProductInfoCommand,
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

	if err := product.UpdateInfo(cmd.Title, cmd.Description); err != nil {
		return err
	}

	return uc.products.Save(ctx, product)
}
