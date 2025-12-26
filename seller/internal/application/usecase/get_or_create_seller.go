package usecase

import (
	aport "clirzy/seller/internal/application/port"
	"clirzy/seller/internal/domain"
	"clirzy/seller/internal/domain/entity"
	"clirzy/seller/internal/domain/port"
	"clirzy/seller/internal/domain/valueobject"
	"context"
)

type GetOrCreateSellerCommand struct {
	UserID string
}

type GetOrCreateSellerUseCase struct {
	sellers port.SellersRepo
	users   aport.UserService
}

func NewGetOrCreateSellerUseCase(
	sellers port.SellersRepo,
	users aport.UserService,
) *GetOrCreateSellerUseCase {
	return &GetOrCreateSellerUseCase{
		sellers: sellers,
		users:   users,
	}
}

func (uc *GetOrCreateSellerUseCase) Execute(ctx context.Context, cmd GetOrCreateSellerCommand) (*entity.Seller, error) {
	userID, err := valueobject.ToUserID(cmd.UserID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}

	user, err := uc.users.FindByID(ctx, string(userID))
	if err != nil || user == nil {
		return nil, domain.ErrNoUserFound
	}

	seller, err := uc.sellers.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if seller == nil {
		seller = entity.NewSeller(valueobject.NewSellerID(), userID)

		if err := uc.sellers.Save(ctx, seller); err != nil {
			return nil, err
		}
	}

	return seller, nil
}
