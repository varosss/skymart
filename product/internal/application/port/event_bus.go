package port

import (
	"clirzy/product/internal/domain/event"
	"context"
)

type EventHandler interface {
	Handle(ctx context.Context, event event.Event) error
}

type EventBus interface {
	Publish(ctx context.Context, events []event.Event) error
	Subscribe(eventType string, handler EventHandler) error
	Start(ctx context.Context) error
	Close() error
}
