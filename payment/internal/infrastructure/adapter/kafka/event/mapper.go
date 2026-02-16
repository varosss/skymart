package event

import (
	domevent "clirzy/payment/internal/domain/event"
	"clirzy/payment/internal/domain/valueobject"
	"clirzy/payment/internal/infrastructure/adapter/kafka/payload"
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

	case *domevent.PaymentCreated:
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

	case *domevent.PaymentSucceeded:
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

	case *domevent.PaymentCanceled:
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

func (m *EventMapper) FromMessage(
	msg pkgkafka.Message,
) (domevent.Event, error) {
	switch msg.EventType {

	case "payment.created":
		data, ok := msg.Payload.(payload.PaymentCreatedPayload)
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

		return domevent.PaymentCreatedFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.PaymentID(data.PaymentID),
			valueobject.InvoiceID(data.InvoiceID),
			money,
			occurredAt,
		), nil

	case "payment.succeeded":
		data, ok := msg.Payload.(payload.PaymentSucceededPayload)
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

		return domevent.PaymentSucceededFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.PaymentID(data.PaymentID),
			valueobject.InvoiceID(data.InvoiceID),
			money,
			occurredAt,
		), nil

	case "payment.canceled":
		data, ok := msg.Payload.(payload.PaymentCanceledPayload)
		if !ok {
			return nil, fmt.Errorf("invalid payload for event %s", msg.EventType)
		}

		occurredAt, err := time.Parse(time.RFC3339, msg.OccurredAt)
		if err != nil {
			return nil, err
		}

		return domevent.PaymentCanceledFromPrimitives(
			valueobject.EventID(msg.EventID),
			valueobject.PaymentID(data.PaymentID),
			valueobject.InvoiceID(data.InvoiceID),
			occurredAt,
		), nil

	default:
		return nil, errors.New("unknown event type")
	}
}
