package entity

import (
	"clirzy/order/internal/domain"
	"clirzy/order/internal/domain/event"
	"clirzy/order/internal/domain/valueobject"
)

type Order struct {
	id      valueobject.OrderID
	buyerID valueobject.BuyerID
	status  valueobject.Status
	items   []OrderItem

	events []event.Event
}

func NewOrder(
	buyerID valueobject.BuyerID,
	items []OrderItem,
) (*Order, error) {

	if len(items) == 0 {
		return nil, domain.ErrEmptyOrder
	}

	o := &Order{
		id:      valueobject.NewOrderID(),
		buyerID: buyerID,
		items:   items,
		status:  valueobject.StatusDraft,
	}

	snapshots := make([]event.OrderItemSnapshot, len(items))
	for i, item := range items {
		snapshots[i] = event.OrderItemSnapshot{
			ProductID: string(item.productID),
			SellerID:  string(item.sellerID),
			Price:     item.price.Amount,
			Qty:       item.qty,
		}
	}

	o.addEvent(event.NewOrderCreated(o.id, o.buyerID, snapshots))

	return o, nil
}

func (o *Order) ID() valueobject.OrderID {
	return o.id
}

func (o *Order) addEvent(e event.Event) {
	o.events = append(o.events, e)
}

func (o *Order) PullEvents() []event.Event {
	events := o.events
	o.events = nil
	return events
}
