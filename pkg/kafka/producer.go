// internal/kafka/producer.go
package kafka

import (
	"context"
	"encoding/json"

	kafkago "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafkago.Writer
}

func NewProducer(cfg Config) *Producer {
	return &Producer{
		writer: &kafkago.Writer{
			Addr:         kafkago.TCP(cfg.Brokers...),
			Topic:        cfg.Topic,
			Balancer:     &kafkago.LeastBytes{},
			RequiredAcks: kafkago.RequireOne,
		},
	}
}

func (p *Producer) Send(ctx context.Context, key string, value []byte) error {
	return p.writer.WriteMessages(ctx, kafkago.Message{
		Key:   []byte(key),
		Value: value,
	})
}

func (p *Producer) SendBatch(ctx context.Context, msgs []*Message) error {
	kafkaMessages := make([]kafkago.Message, len(msgs))

	for i, m := range msgs {
		value, err := json.Marshal(m)
		if err != nil {
			return err
		}

		kafkaMessages[i] = kafkago.Message{
			Key:   []byte(m.AggregateID),
			Value: value,
		}
	}

	return p.writer.WriteMessages(ctx, kafkaMessages...)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
