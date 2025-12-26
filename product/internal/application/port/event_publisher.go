package port

import "clirzy/product/internal/domain/event"

type EventPublisher interface {
	Publish(events []event.Event) error
}
