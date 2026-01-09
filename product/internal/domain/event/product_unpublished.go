package event

import (
	"clirzy/product/internal/domain/valueobject"
)

type ProductUnpublished struct {
	ProductID valueobject.ProductID
	SellerID  valueobject.SellerID
}

func NewProductUnpublished(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
) ProductUnpublished {
	return ProductUnpublished{
		ProductID: productID,
		SellerID:  sellerID,
	}
}

func (e ProductUnpublished) Name() string {
	return "product.unpublished"
}
