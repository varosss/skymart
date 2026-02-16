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
	total   valueobject.Money
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

	var totalAmount int64
	var currency string

	snapshots := make([]*event.OrderItemSnapshot, len(items))
	for i, item := range items {
		if i == 0 {
			currency = item.price.Currency()
		}

		if item.price.Currency() != currency {
			return nil, domain.ErrMixedCurrencies
		}

		lineTotal := item.price.Amount() * item.qty
		totalAmount += lineTotal

		snapshots[i] = event.NewOrderItemSnapshot(
			item.productID,
			item.sellerID,
			item.price.Amount(),
			item.price.Currency(),
			item.qty,
		)
	}

	total, err := valueobject.NewMoney(totalAmount, currency)
	if err != nil {
		return nil, err
	}

	o := &Order{
		id:      valueobject.NewOrderID(),
		buyerID: buyerID,
		status:  valueobject.StatusDraft,
		total:   total,
		items:   items,
	}

	o.addEvent(event.NewOrderCreated(o.id, o.buyerID, total, snapshots))

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
