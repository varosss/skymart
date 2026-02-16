package usecase

import (
	aport "clirzy/product/internal/application/port"
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/port"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type UpdateProductCommand struct {
	UserID      valueobject.UserID
	ProductID   valueobject.ProductID
	Title       *string
	Description *string
	Price       *int64
	Currency    *string
	Status      *string
}

type UpdateProductUseCase struct {
	sellers  aport.SellerGateway
	products port.ProductRepo
	bus      aport.EventBus
}

func NewUpdateProductUseCase(
	sellers aport.SellerGateway,
	products port.ProductRepo,
	bus aport.EventBus,
) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		sellers:  sellers,
		products: products,
		bus:      bus,
	}
}

func (uc *UpdateProductUseCase) Execute(ctx context.Context, cmd UpdateProductCommand) error {
	seller, err := uc.sellers.GetByUserID(ctx, cmd.UserID)
	if err != nil {
		return err
	}
	if !seller.IsActive {
		return domain.ErrSellerInactive
	}

	product, err := uc.products.GetByID(ctx, cmd.ProductID)
	if err != nil {
		return err
	}

	if product.SellerID() != seller.ID {
		return domain.ErrProductNotOwnedBySeller
	}

	if cmd.Status != nil {
		switch valueobject.Status(*cmd.Status) {
		case valueobject.StatusPublished:
			if err := product.Publish(); err != nil {
				return err
			}
		case valueobject.StatusUnpublished:
			if err := product.Unpublish(); err != nil {
				return err
			}
		case valueobject.StatusArchived:
			if err := product.Archive(); err != nil {
				return err
			}
		}
	}

	if cmd.Title != nil || cmd.Description != nil {
		title := product.Title()
		description := product.Description()

		if cmd.Title != nil {
			title = *cmd.Title
		}
		if cmd.Description != nil {
			description = *cmd.Description
		}

		if err := product.UpdateInfo(title, description); err != nil {
			return err
		}
	}

	if cmd.Price != nil && cmd.Currency != nil {
		price, err := valueobject.NewMoney(*cmd.Price, *cmd.Currency)
		if err != nil {
			return err
		}

		if err := product.UpdatePrice(price); err != nil {
			return err
		}
	}

	if err := uc.products.Save(ctx, product); err != nil {
		return err
	}

	return uc.bus.Publish(ctx, product.PullEvents())
}
