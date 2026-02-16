package app

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, app *Application) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, len(app.components))

	for _, c := range app.components {
		go func(comp Component) {
			log.Println("starting", comp.Name())
			if err := comp.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
				errCh <- err
			}
		}(c)
	}

	select {
	case <-ctx.Done():
		log.Println("shutting down components...")

		shutdownCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for i := len(app.components) - 1; i >= 0; i-- {
			comp := app.components[i]
			log.Println("stopping", comp.Name())
			_ = comp.Shutdown(shutdownCtx)
		}

		for i := len(app.closers) - 1; i >= 0; i-- {
			_ = app.closers[i]()
		}

		return nil

	case err := <-errCh:
		return err
	}
}
