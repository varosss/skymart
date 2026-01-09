package app

import "context"

type Component interface {
	Name() string
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
