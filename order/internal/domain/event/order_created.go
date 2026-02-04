package event

import (
	"clirzy/order/internal/domain/valueobject"
	"time"
)

type OrderCreated struct {
	eventID    valueobject.EventID
	orderID    valueobject.OrderID
	buyerID    valueobject.BuyerID
	total      valueobject.Money
	items      []OrderItemSnapshot
	occurredAt time.Time
}

func NewOrderCreated(
	orderID valueobject.OrderID,
	buyerID valueobject.BuyerID,
	total valueobject.Money,
	items []OrderItemSnapshot,
) OrderCreated {
	return OrderCreated{
		eventID:    valueobject.NewEventID(),
		orderID:    orderID,
		buyerID:    buyerID,
		total:      total,
		items:      items,
		occurredAt: time.Now(),
	}
}

func (e OrderCreated) ID() string {
	return e.eventID.String()
}

func (OrderCreated) Type() string {
	return "order.created"
}

func (e OrderCreated) AggregateID() string {
	return e.orderID.String()
}

func (OrderCreated) AggregateType() string {
	return "order"
}

func (e OrderCreated) OccurredAt() time.Time {
	return e.occurredAt
}

func (e OrderCreated) OrderID() valueobject.OrderID {
	return e.orderID
}

func (e OrderCreated) BuyerID() valueobject.BuyerID {
	return e.buyerID
}

func (e OrderCreated) Total() valueobject.Money {
	return e.total
}

func (e OrderCreated) Items() []OrderItemSnapshot {
	return e.items
}
