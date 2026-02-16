package valueobject

import "github.com/google/uuid"

type EventID string

func NewEventID() EventID {
	return EventID(uuid.New().String())
}

func (id EventID) String() string {
	return string(id)
}
