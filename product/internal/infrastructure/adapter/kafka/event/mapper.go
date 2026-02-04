package event

import (
	pkgkafka "clirzy/pkg/kafka"
	domevent "clirzy/product/internal/domain/event"
	"clirzy/product/internal/infrastructure/adapter/kafka/payload"
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

	case domevent.ProductCreated:
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

	case domevent.ProductArchived:
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

	case domevent.ProductPublished:
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

	case domevent.ProductUnpublished:
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

	case domevent.ProductInfoUpdated:
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

	case domevent.ProductPriceUpdated:
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
