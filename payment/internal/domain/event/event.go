package event

import "time"

type Event interface {
	ID() string
	Type() string
	AggregateID() string
	AggregateType() string
	OccurredAt() time.Time
}
