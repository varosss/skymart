package event

import (
	domevent "clirzy/order/internal/domain/event"
	"clirzy/order/internal/domain/valueobject"
	"clirzy/order/internal/infrastructure/adapter/kafka/payload"
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

	case *domevent.OrderCreated:
		items := make([]payload.OrderCreatedItem, len(ev.Items()))

		for i, item := range ev.Items() {
			items[i] = payload.OrderCreatedItem{
				ProductID: item.ProductID().String(),
				SellerID:  item.SellerID().String(),
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

func (m *EventMapper) FromMessage(
	msg pkgkafka.Message,
) (domevent.Event, error) {
	switch msg.EventType {

	case "order.created":
		data, ok := msg.Payload.(payload.OrderCreatedPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		items := make([]*domevent.OrderItemSnapshot, len(data.Items))

		for i, item := range data.Items {
			items[i] = domevent.NewOrderItemSnapshot(
				valueobject.ProductID(item.ProductID),
				valueobject.SellerID(item.SellerID),
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

		return domevent.OrderCreatedFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.OrderID(data.OrderID),
			valueobject.BuyerID(data.BuyerID),
			money,
			items,
			occurredAt,
		), nil

	default:
		return nil, errors.New("unknown event type")
	}
}
