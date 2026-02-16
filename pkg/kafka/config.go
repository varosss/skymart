package kafka

import "time"

type Config struct {
	Brokers []string
	Topic   string
	GroupID string
	Timeout time.Duration
}
