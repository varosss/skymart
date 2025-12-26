package event

import (
	"clirzy/product/internal/domain/valueobject"
)

type ProductPublished struct {
	ProductID valueobject.ProductID
	SellerID  valueobject.SellerID
}

func NewProductPublished(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
) ProductPublished {
	return ProductPublished{
		ProductID: productID,
		SellerID:  sellerID,
	}
}

func (e ProductPublished) Name() string {
	return "product.published"
}
