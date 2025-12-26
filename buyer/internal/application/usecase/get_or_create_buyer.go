package usecase

import (
	aport "clirzy/buyer/internal/application/port"
	"clirzy/buyer/internal/domain"
	"clirzy/buyer/internal/domain/entity"
	"clirzy/buyer/internal/domain/port"
	"clirzy/buyer/internal/domain/valueobject"
	"context"
)

type GetOrCreateBuyerCommand struct {
	UserID string
}

type GetOrCreateBuyerUseCase struct {
	buyers port.BuyersRepo
	users  aport.UserService
}

func NewGetOrCreateBuyerUseCase(
	buyers port.BuyersRepo,
	users aport.UserService,
) *GetOrCreateBuyerUseCase {
	return &GetOrCreateBuyerUseCase{
		buyers: buyers,
		users:  users,
	}
}

func (uc *GetOrCreateBuyerUseCase) Execute(ctx context.Context, cmd GetOrCreateBuyerCommand) (*entity.Buyer, error) {
	userID, err := valueobject.ToUserID(cmd.UserID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}

	user, err := uc.users.FindByID(ctx, string(userID))
	if err != nil || user == nil {
		return nil, domain.ErrNoUserFound
	}

	buyer, err := uc.buyers.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if buyer == nil {
		buyer = entity.NewBuyer(valueobject.NewBuyerID(), userID)

		if err := uc.buyers.Save(ctx, buyer); err != nil {
			return nil, err
		}
	}

	return buyer, nil
}
