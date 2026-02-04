package event

import (
	pkgkafka "clirzy/pkg/kafka"
	domevent "clirzy/product/internal/domain/event"
	"context"
)

type KafkaEventBus struct {
	mapper   *EventMapper
	producer *pkgkafka.Producer
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
