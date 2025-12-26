package usecase

import (
	"clirzy/product/internal/application/service"
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/port"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type UpdateProductPriceCommand struct {
	SellerID  string
	ProductID string
	Price     int64
	Currency  string
}

type UpdateProductPriceUseCase struct {
	access   *service.ProductAccessService
	products port.ProductsRepo
}

func NewUpdateProductPriceUseCase(
	access *service.ProductAccessService,
	products port.ProductsRepo,
) *UpdateProductPriceUseCase {
	return &UpdateProductPriceUseCase{access, products}
}

func (uc *UpdateProductPriceUseCase) Execute(
	ctx context.Context,
	cmd UpdateProductPriceCommand,
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

	price := valueobject.NewMoney(cmd.Price, cmd.Currency)
	if err := product.UpdatePrice(price); err != nil {
		return err
	}

	return uc.products.Save(ctx, product)
}
