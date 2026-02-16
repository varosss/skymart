package usecase

import (
	"clirzy/seller/internal/domain/entity"
	"clirzy/seller/internal/domain/port"
	"clirzy/seller/internal/domain/valueobject"
	"context"
	"errors"
)

type CreateSellerCommand struct {
	UserID valueobject.UserID
}

type CreateSellerUseCase struct {
	sellers port.SellerRepo
}

func NewCreateSellerUseCase(
	sellers port.SellerRepo,
) *CreateSellerUseCase {
	return &CreateSellerUseCase{
		sellers: sellers,
	}
}

func (uc *CreateSellerUseCase) Execute(ctx context.Context, cmd CreateSellerCommand) (*entity.Seller, error) {
	seller, err := uc.sellers.FindByUserID(ctx, cmd.UserID)
	if err != nil {
		return nil, err
	}

	if seller != nil {
		return nil, errors.New("seller already exists")
	}

	seller = entity.NewSeller(cmd.UserID)
	if err := uc.sellers.Save(ctx, seller); err != nil {
		return nil, err
	}

	return seller, nil
}
