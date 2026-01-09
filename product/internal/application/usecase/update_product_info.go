package usecase

import (
	aport "clirzy/product/internal/application/port"
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
	access    *service.ProductAccessService
	products  port.ProductsRepo
	publisher aport.EventPublisher
}

func NewUpdateProductInfoUseCase(
	access *service.ProductAccessService,
	products port.ProductsRepo,
	publisher aport.EventPublisher,
) *UpdateProductInfoUseCase {
	return &UpdateProductInfoUseCase{
		access:    access,
		products:  products,
		publisher: publisher,
	}
}

func (uc *UpdateProductInfoUseCase) Execute(
	ctx context.Context,
	cmd UpdateProductInfoCommand,
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

	if err := product.UpdateInfo(cmd.Title, cmd.Description); err != nil {
		return err
	}

	if err := uc.products.Save(ctx, product); err != nil {
		return err
	}

	return uc.publisher.Publish(product.PullEvents())
}
