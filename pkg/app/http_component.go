package app

import (
	"clirzy/pkg/httpserver"
	"context"
)

type HTTPComponent struct {
	srv *httpserver.Server
}

func NewHTTPComponent(srv *httpserver.Server) *HTTPComponent {
	return &HTTPComponent{srv: srv}
}

func (c *HTTPComponent) Name() string {
	return "http-server"
}

func (c *HTTPComponent) Run(ctx context.Context) error {
	return c.srv.Run()
}

func (c *HTTPComponent) Shutdown(ctx context.Context) error {
	return c.srv.Shutdown(ctx)
}
