package usecase

import (
	aport "clirzy/buyer/internal/application/port"
	"clirzy/buyer/internal/domain"
	"clirzy/buyer/internal/domain/entity"
	"clirzy/buyer/internal/domain/port"
	"clirzy/buyer/internal/domain/valueobject"
	"context"
	"errors"
)

type CreateBuyerCommand struct {
	UserID string
}

type CreateBuyerUseCase struct {
	buyers port.BuyerRepo
}

func NewCreateBuyerUseCase(
	buyers port.BuyerRepo,
	users aport.UserGateway,
) *CreateBuyerUseCase {
	return &CreateBuyerUseCase{
		buyers: buyers,
	}
}

func (uc *CreateBuyerUseCase) Execute(ctx context.Context, cmd CreateBuyerCommand) (*entity.Buyer, error) {
	userID, err := valueobject.ParseUserID(cmd.UserID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}

	buyer, err := uc.buyers.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if buyer != nil {
		return nil, errors.New("buyer already exists")
	}

	buyer = entity.NewBuyer(userID)
	if err := uc.buyers.Save(ctx, buyer); err != nil {
		return nil, err
	}

	return buyer, nil
}
