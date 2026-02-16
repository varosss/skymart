package event

import (
	pkgkafka "clirzy/pkg/kafka"
	domevent "clirzy/product/internal/domain/event"
	"clirzy/product/internal/domain/valueobject"
	"clirzy/product/internal/infrastructure/adapter/kafka/payload"
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

	case *domevent.ProductCreated:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.ProductCreatedPayload{
				ProductID: ev.ProductID().String(),
				SellerID:  ev.SellerID().String(),
			},
		}, nil

	case *domevent.ProductArchived:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.ProductArchivedPayload{
				ProductID: ev.ProductID().String(),
				SellerID:  ev.SellerID().String(),
			},
		}, nil

	case *domevent.ProductPublished:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.ProductPublishedPayload{
				ProductID: ev.ProductID().String(),
				SellerID:  ev.SellerID().String(),
			},
		}, nil

	case *domevent.ProductUnpublished:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.ProductUnpublishedPayload{
				ProductID: ev.ProductID().String(),
				SellerID:  ev.SellerID().String(),
			},
		}, nil

	case *domevent.ProductInfoUpdated:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.ProductInfoUpdatedPayload{
				ProductID:   ev.ProductID().String(),
				SellerID:    ev.SellerID().String(),
				Title:       ev.Title(),
				Description: ev.Description(),
			},
		}, nil

	case *domevent.ProductPriceUpdated:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.ProductPriceUpdatedPayload{
				ProductID: ev.ProductID().String(),
				SellerID:  ev.SellerID().String(),
				Amount:    ev.Price().Amount(),
				Currency:  ev.Price().Currency(),
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

	case "product.created":
		data, ok := msg.Payload.(payload.ProductCreatedPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		occurredAt, err := time.Parse(time.RFC3339, msg.OccurredAt)
		if err != nil {
			return nil, err
		}

		return domevent.ProductCreatedFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.ProductID(data.ProductID),
			valueobject.SellerID(data.SellerID),
			occurredAt,
		), nil

	case "product.archived":
		data, ok := msg.Payload.(payload.ProductArchivedPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		occurredAt, err := time.Parse(time.RFC3339, msg.OccurredAt)
		if err != nil {
			return nil, err
		}

		return domevent.ProductArchivedFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.ProductID(data.ProductID),
			valueobject.SellerID(data.SellerID),
			occurredAt,
		), nil

	case "product.published":
		data, ok := msg.Payload.(payload.ProductPublishedPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		occurredAt, err := time.Parse(time.RFC3339, msg.OccurredAt)
		if err != nil {
			return nil, err
		}

		return domevent.ProductPublishedFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.ProductID(data.ProductID),
			valueobject.SellerID(data.SellerID),
			occurredAt,
		), nil

	case "product.unpublished":
		data, ok := msg.Payload.(payload.ProductUnpublishedPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		occurredAt, err := time.Parse(time.RFC3339, msg.OccurredAt)
		if err != nil {
			return nil, err
		}

		return domevent.ProductUnpublishedFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.ProductID(data.ProductID),
			valueobject.SellerID(data.SellerID),
			occurredAt,
		), nil

	case "product.price_updated":
		data, ok := msg.Payload.(payload.ProductPriceUpdatedPayload)
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

		return domevent.ProductPriceUpdatedFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.ProductID(data.ProductID),
			valueobject.SellerID(data.SellerID),
			money,
			occurredAt,
		), nil

	default:
		return nil, errors.New("unknown event type")
	}
}
