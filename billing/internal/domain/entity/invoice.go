package entity

import (
	"clirzy/billing/internal/domain"
	"clirzy/billing/internal/domain/event"
	"clirzy/billing/internal/domain/valueobject"
)

type Invoice struct {
	id      valueobject.InvoiceID
	orderID valueobject.OrderID
	buyerID valueobject.BuyerID
	status  valueobject.InvoiceStatus
	total   valueobject.Money
	items   []InvoiceItem

	events []event.Event
}

func NewInvoice(
	buyerID valueobject.BuyerID,
	orderID valueobject.OrderID,
	items []InvoiceItem,
) (*Invoice, error) {

	if len(items) == 0 {
		return nil, domain.ErrEmptyInvoice
	}

	var totalAmount int64
	var currency string

	snapshots := make([]*event.InvoiceItemSnapshot, len(items))

	for i, item := range items {
		if i == 0 {
			currency = item.price.Currency()
		}

		if item.price.Currency() != currency {
			return nil, domain.ErrMixedCurrencies
		}

		lineTotal := item.price.Amount() * item.qty
		totalAmount += lineTotal

		snapshots[i] = event.NewInvoiceItemSnapshot(
			item.productID,
			item.price.Amount(),
			item.price.Currency(),
			item.qty,
		)
	}

	total, err := valueobject.NewMoney(totalAmount, currency)
	if err != nil {
		return nil, err
	}

	inv := &Invoice{
		id:      valueobject.NewInvoiceID(),
		buyerID: buyerID,
		orderID: orderID,
		items:   items,
		total:   total,
		status:  valueobject.InvoiceStatusPending,
	}

	inv.addEvent(event.NewInvoiceCreated(
		inv.id,
		inv.buyerID,
		total,
		snapshots,
	))

	return inv, nil
}

func (i *Invoice) ID() valueobject.InvoiceID {
	return i.id
}

func (i *Invoice) Status() valueobject.InvoiceStatus {
	return i.status
}

func (i *Invoice) Total() valueobject.Money {
	return i.total
}

func (i *Invoice) MarkPaid() error {
	if i.status != valueobject.InvoiceStatusPending {
		return domain.ErrInvalidInvoiceState
	}

	i.status = valueobject.InvoiceStatusPaid

	i.addEvent(event.NewInvoicePaid(i.id, i.buyerID, i.total))

	return nil
}

func (i *Invoice) Cancel() error {
	if i.status == valueobject.InvoiceStatusPaid {
		return domain.ErrCannotCancelPaidInvoice
	}

	i.status = valueobject.InvoiceStatusCanceled

	i.addEvent(event.NewInvoiceCanceled(i.id, i.buyerID))

	return nil
}

func (i *Invoice) addEvent(e event.Event) {
	i.events = append(i.events, e)
}

func (i *Invoice) PullEvents() []event.Event {
	events := i.events
	i.events = nil
	return events
}
