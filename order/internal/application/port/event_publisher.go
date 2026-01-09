package port

import "clirzy/order/internal/domain/event"

type EventPublisher interface {
	Publish(events []event.Event) error
}
