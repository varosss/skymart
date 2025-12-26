package service

import (
	aport "clirzy/product/internal/application/port"
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/entity"
	"clirzy/product/internal/domain/port"
	"clirzy/product/internal/domain/valueobject"
	"context"
)

type ProductAccessService struct {
	sellers  aport.SellerService
	products port.ProductsRepo
}

func NewProductAccessService(
	sellers aport.SellerService,
	products port.ProductsRepo,
) *ProductAccessService {
	return &ProductAccessService{sellers, products}
}

func (s *ProductAccessService) LoadForSeller(
	ctx context.Context,
	sellerID valueobject.SellerID,
	productID valueobject.ProductID,
) (*entity.Product, error) {

	exists, err := s.sellers.Exists(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.ErrSellerNotFound
	}

	active, err := s.sellers.IsActive(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	if !active {
		return nil, domain.ErrSellerInactive
	}

	product, err := s.products.GetByID(ctx, productID)
	if err != nil {
		return nil, domain.ErrProductNotFound
	}

	if product.SellerID() != sellerID {
		return nil, domain.ErrProductNotOwnedBySeller
	}

	return product, nil
}
