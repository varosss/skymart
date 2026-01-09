package event

import (
	"clirzy/product/internal/domain/valueobject"
)

type ProductArchived struct {
	ProductID valueobject.ProductID
	SellerID  valueobject.SellerID
}

func NewProductArchived(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
) ProductArchived {
	return ProductArchived{
		ProductID: productID,
		SellerID:  sellerID,
	}
}

func (e ProductArchived) Name() string {
	return "product.archived"
}
