package event

import (
	domevent "clirzy/order/internal/domain/event"
	"clirzy/order/internal/infrastructure/adapter/kafka/payload"
	pkgkafka "clirzy/pkg/kafka"
	"errors"
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

	case domevent.OrderCreated:
		items := make([]payload.OrderCreatedItem, len(ev.Items()))

		for i, item := range ev.Items() {
			items[i] = payload.OrderCreatedItem{
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
			Payload: payload.OrderCreatedPayload{
				OrderID:  ev.OrderID().String(),
				BuyerID:  ev.BuyerID().String(),
				Total:    ev.Total().Amount(),
				Currency: ev.Total().Currency(),
				Items:    items,
			},
		}, nil

	default:
		return nil, errors.New("unknown domain event")
	}
}
