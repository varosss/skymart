package kafka

import (
	domevent "clirzy/billing/internal/domain/event"
	"clirzy/billing/internal/domain/valueobject"
	"clirzy/billing/internal/infrastructure/adapter/kafka/payload"
	pkgkafka "clirzy/pkg/kafka"
	"errors"
	"fmt"
	"time"
)

type EventMapper struct{}

func NewEventMapper() *EventMapper {
	return &EventMapper{}
}

func (m *EventMapper) ToMessage(
	e domevent.Event,
) (*pkgkafka.Message, error) {

	switch ev := e.(type) {

	case *domevent.InvoiceCreated:
		items := make([]payload.InvoiceCreatedItem, len(ev.Items()))

		for i, item := range ev.Items() {
			items[i] = payload.InvoiceCreatedItem{
				ProductID: item.ProductID().String(),
				Amount:    item.Amount(),
				Currency:  item.Currency(),
				Qty:       item.Qty(),
			}
		}

		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.InvoiceCreatedPayload{
				InvoiceID: ev.InvoiceID().String(),
				BuyerID:   ev.BuyerID().String(),
				Total:     ev.Total(),
				Currency:  ev.Currency(),
				Items:     items,
			},
		}, nil

	case *domevent.InvoicePaid:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.InvoicePaidPayload{
				InvoiceID: ev.InvoiceID().String(),
				BuyerID:   ev.BuyerID().String(),
				Amount:    ev.Amount(),
				Currency:  ev.Currency(),
			},
		}, nil

	case *domevent.InvoiceCanceled:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.InvoiceCanceledPayload{
				InvoiceID: ev.InvoiceID().String(),
				BuyerID:   ev.BuyerID().String(),
			},
		}, nil

	default:
		return nil, errors.New("unknown domain event")
	}
}

func (m *EventMapper) FromMessage(
	msg pkgkafka.Message,
) (domevent.Event, error) {
	switch msg.EventType {

	case "invoice.created":
		data, ok := msg.Payload.(payload.InvoiceCreatedPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		items := make([]*domevent.InvoiceItemSnapshot, len(data.Items))

		for i, item := range data.Items {
			items[i] = domevent.NewInvoiceItemSnapshot(
				valueobject.ProductID(item.ProductID),
				item.Amount,
				item.Currency,
				item.Qty,
			)
		}

		money, err := valueobject.NewMoney(data.Total, data.Currency)
		if err != nil {
			return nil, err
		}

		occurredAt, err := time.Parse(time.RFC3339, msg.OccurredAt)
		if err != nil {
			return nil, err
		}

		return domevent.InvoiceCreatedFromPrimitives(
			valueobject.EventID(msg.AggregateID),
			valueobject.InvoiceID(data.InvoiceID),
			valueobject.BuyerID(data.BuyerID),
			money,
			items,
			occurredAt,
		), nil

	case "invoice.paid":
		data, ok := msg.Payload.(payload.InvoicePaidPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		money, err := valueobject.NewMoney(data.Amount, data.Currency)
		if err != nil {
			return nil, err
		}

		occurredAt, err := time.Parse(time.RFC3339, msg.OccurredAt)
		if err != nil {
			return nil, err
		}

		return domevent.InvoicePaidFromPrimitives(
			valueobject.EventID(msg.AggregateID),
			valueobject.InvoiceID(data.InvoiceID),
			valueobject.BuyerID(data.BuyerID),
			money,
			occurredAt,
		), nil

	case "invoice.canceled":
		data, ok := msg.Payload.(payload.InvoiceCanceledPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		occurredAt, err := time.Parse(time.RFC3339, msg.OccurredAt)
		if err != nil {
			return nil, err
		}

		return domevent.InvoiceCanceledFromPrimitives(
			valueobject.EventID(msg.AggregateID),
			valueobject.InvoiceID(data.InvoiceID),
			valueobject.BuyerID(data.BuyerID),
			occurredAt,
		), nil

	default:
		return nil, errors.New("unknown event type")
	}
}
