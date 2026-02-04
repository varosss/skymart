package event

import (
	"clirzy/order/internal/domain/valueobject"
)

type OrderItemSnapshot struct {
	productID valueobject.ProductID
	sellerID  valueobject.SellerID
	amount    int64
	currency  string
	qty       int64
}

func NewOrderItemSnapshot(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	amount int64,
	currency string,
	qty int64,
) OrderItemSnapshot {
	return OrderItemSnapshot{
		productID: productID,
		sellerID:  sellerID,
		amount:    amount,
		currency:  currency,
		qty:       qty,
	}
}

func (s OrderItemSnapshot) ProductID() valueobject.ProductID {
	return s.productID
}

func (s OrderItemSnapshot) SellerID() valueobject.SellerID {
	return s.sellerID
}

func (s OrderItemSnapshot) Amount() int64 {
	return s.amount
}

func (s OrderItemSnapshot) Currency() string {
	return s.currency
}

func (s OrderItemSnapshot) Qty() int64 {
	return s.qty
}
