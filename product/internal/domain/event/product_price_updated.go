package event

import (
	"clirzy/product/internal/domain/valueobject"
)

type ProductPriceUpdated struct {
	ProductID valueobject.ProductID
	SellerID  valueobject.SellerID
	Price     valueobject.Money
}

func NewProductPriceUpdated(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	price valueobject.Money,
) ProductPriceUpdated {
	return ProductPriceUpdated{
		ProductID: productID,
		SellerID:  sellerID,
		Price:     price,
	}
}

func (e ProductPriceUpdated) Name() string {
	return "product.price_updated"
}
