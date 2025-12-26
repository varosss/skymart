package event

import (
	"clirzy/product/internal/domain/valueobject"
)

type ProductCreated struct {
	ProductID valueobject.ProductID
	SellerID  valueobject.SellerID
}

func NewProductCreated(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
) ProductCreated {
	return ProductCreated{
		ProductID: productID,
		SellerID:  sellerID,
	}
}

func (e ProductCreated) Name() string {
	return "product.created"
}
