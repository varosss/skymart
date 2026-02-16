package kafka

type Message struct {
	EventID       string      `json:"event_id"`
	EventType     string      `json:"event_type"`
	AggregateID   string      `json:"aggregate_id"`
	AggregateType string      `json:"aggregate_type"`
	OccurredAt    string      `json:"occurred_at"`
	Payload       interface{} `json:"payload"`
}
