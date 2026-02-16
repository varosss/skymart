package event

import (
	aport "clirzy/order/internal/application/port"
	domevent "clirzy/order/internal/domain/event"
	pkgkafka "clirzy/pkg/kafka"
	"context"
	"encoding/json"
	"fmt"
)

type KafkaEventBus struct {
	mapper *EventMapper

	producer *pkgkafka.Producer
	consumer *pkgkafka.Consumer

	handlers map[string]aport.EventHandler
}

func (p *KafkaEventBus) Publish(ctx context.Context, events []domevent.Event) error {
	messages := make([]*pkgkafka.Message, len(events))

	for i, e := range events {
		msg, err := p.mapper.ToMessage(e)
		if err != nil {
			return err
		}

		messages[i] = msg
	}

	p.producer.SendBatch(ctx, messages)

	return nil
}

func (b *KafkaEventBus) Subscribe(eventType string, handler aport.EventHandler) error {
	if _, exists := b.handlers[eventType]; exists {
		return fmt.Errorf("handler already registered for %s", eventType)
	}

	b.handlers[eventType] = handler
	return nil
}

func (b *KafkaEventBus) Start(ctx context.Context) error {
	return b.consumer.Start(ctx, b.handle)
}

func (b *KafkaEventBus) handle(value []byte) error {
	var msg pkgkafka.Message
	if err := json.Unmarshal(value, &msg); err != nil {
		return err
	}

	handler, ok := b.handlers[msg.EventType]
	if !ok {
		return nil
	}

	event, err := b.mapper.FromMessage(msg)
	if err != nil {
		return err
	}

	return handler.Handle(context.Background(), event)
}
