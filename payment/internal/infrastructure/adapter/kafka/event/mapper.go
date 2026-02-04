package event

import (
	domevent "clirzy/payment/internal/domain/event"
	"clirzy/payment/internal/infrastructure/adapter/kafka/payload"
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

	case domevent.PaymentCreated:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.PaymentCreatedPayload{
				PaymentID: ev.PaymentID().String(),
				InvoiceID: ev.InvoiceID().String(),
				Amount:    ev.Amount(),
				Currency:  ev.Currency(),
			},
		}, nil

	case domevent.PaymentSucceeded:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.PaymentSucceededPayload{
				PaymentID: ev.PaymentID().String(),
				InvoiceID: ev.InvoiceID().String(),
				Amount:    ev.Amount(),
				Currency:  ev.Currency(),
			},
		}, nil

	case domevent.PaymentCanceled:
		return &pkgkafka.Message{
			EventID:       ev.ID(),
			EventType:     ev.Type(),
			AggregateID:   ev.AggregateID(),
			AggregateType: ev.AggregateType(),
			OccurredAt:    ev.OccurredAt().Format(time.RFC3339),
			Payload: payload.PaymentCanceledPayload{
				PaymentID: ev.PaymentID().String(),
				InvoiceID: ev.InvoiceID().String(),
			},
		}, nil

	default:
		return nil, errors.New("unknown domain event")
	}
}
