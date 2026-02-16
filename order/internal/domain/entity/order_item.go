package entity

import "clirzy/order/internal/domain/valueobject"

type OrderItem struct {
	productID valueobject.ProductID
	sellerID  valueobject.SellerID
	price     valueobject.Money
	qty       int64
}

func NewOrderItem(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	price valueobject.Money,
	qty int64,
) OrderItem {
	return OrderItem{
		productID: productID,
		sellerID:  sellerID,
		price:     price,
		qty:       qty,
	}
}
