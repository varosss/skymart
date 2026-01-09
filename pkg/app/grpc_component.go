package app

import (
	"clirzy/pkg/grpcserver"
	"context"
)

type GRPCComponent struct {
	srv *grpcserver.Server
}

func NewGRPCComponent(srv *grpcserver.Server) *GRPCComponent {
	return &GRPCComponent{
		srv: srv,
	}
}

func (c *GRPCComponent) Name() string {
	return "grpc-server"
}

func (c *GRPCComponent) Run(ctx context.Context) error {
	return c.srv.Run()
}

func (c *GRPCComponent) Shutdown(ctx context.Context) error {
	c.srv.GracefulStop()
	return nil
}
