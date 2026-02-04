package usecase

import (
	aport "clirzy/product/internal/application/port"
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
	products port.ProductRepo
	bus      aport.EventBus
}

func NewUpdateProductPriceUseCase(
	access *service.ProductAccessService,
	products port.ProductRepo,
	bus aport.EventBus,
) *UpdateProductPriceUseCase {
	return &UpdateProductPriceUseCase{
		access:   access,
		products: products,
		bus:      bus,
	}
}

func (uc *UpdateProductPriceUseCase) Execute(
	ctx context.Context,
	cmd UpdateProductPriceCommand,
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

	price := valueobject.NewMoney(cmd.Price, cmd.Currency)
	if err := product.UpdatePrice(price); err != nil {
		return err
	}

	if err := uc.products.Save(ctx, product); err != nil {
		return err
	}

	return uc.bus.Publish(ctx, product.PullEvents())
}
