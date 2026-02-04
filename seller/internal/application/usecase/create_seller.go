package usecase

import (
	"clirzy/seller/internal/domain"
	"clirzy/seller/internal/domain/entity"
	"clirzy/seller/internal/domain/port"
	"clirzy/seller/internal/domain/valueobject"
	"context"
	"errors"
)

type CreateSellerCommand struct {
	UserID string
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
	userID, err := valueobject.ParseUserID(cmd.UserID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}

	seller, err := uc.sellers.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if seller != nil {
		return nil, errors.New("seller already exists")
	}

	seller = entity.NewSeller(userID)
	if err := uc.sellers.Save(ctx, seller); err != nil {
		return nil, err
	}

	return seller, nil
}
