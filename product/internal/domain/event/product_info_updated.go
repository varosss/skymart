package event

import (
	"clirzy/product/internal/domain/valueobject"
)

type ProductInfoUpdated struct {
	ProductID   valueobject.ProductID
	SellerID    valueobject.SellerID
	Title       string
	Description string
}

func NewProductInfoUpdated(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	title string,
	description string,
) ProductInfoUpdated {
	return ProductInfoUpdated{
		ProductID:   productID,
		SellerID:    sellerID,
		Title:       title,
		Description: description,
	}
}

func (e ProductInfoUpdated) Name() string {
	return "product.price_updated"
}
