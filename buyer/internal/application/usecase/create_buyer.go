package usecase

import (
	"clirzy/buyer/internal/domain/entity"
	"clirzy/buyer/internal/domain/port"
	"clirzy/buyer/internal/domain/valueobject"
	"context"
	"errors"
)

type CreateBuyerCommand struct {
	UserID valueobject.UserID
}

type CreateBuyerUseCase struct {
	buyers port.BuyerRepo
}

func NewCreateBuyerUseCase(
	buyers port.BuyerRepo,
) *CreateBuyerUseCase {
	return &CreateBuyerUseCase{
		buyers: buyers,
	}
}

func (uc *CreateBuyerUseCase) Execute(ctx context.Context, cmd CreateBuyerCommand) (*entity.Buyer, error) {
	buyer, err := uc.buyers.FindByUserID(ctx, cmd.UserID)
	if err != nil {
		return nil, err
	}

	if buyer != nil {
		return nil, errors.New("buyer already exists")
	}

	buyer = entity.NewBuyer(cmd.UserID)
	if err := uc.buyers.Save(ctx, buyer); err != nil {
		return nil, err
	}

	return buyer, nil
}
