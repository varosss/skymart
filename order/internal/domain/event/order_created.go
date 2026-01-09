package event

import (
	"clirzy/order/internal/domain/valueobject"
)

type OrderCreated struct {
	OrderID valueobject.OrderID
	BuyerID valueobject.BuyerID
	Items   []OrderItemSnapshot
}

func NewOrderCreated(
	orderID valueobject.OrderID,
	buyerID valueobject.BuyerID,
	items []OrderItemSnapshot,
) OrderCreated {
	return OrderCreated{
		OrderID: orderID,
		BuyerID: buyerID,
		Items:   items,
	}
}

func (e OrderCreated) Name() string {
	return "order.created"
}
