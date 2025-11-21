package bootstrap

import (
	"log"
	"sync"
)

type HTTPServer interface {
	Run() error
}

type GRPCServer interface {
	Run() error
}

type Server struct {
	httpServer HTTPServer
	grpcServer GRPCServer
}

type Option func(*Server)

func WithHTTP(server HTTPServer) Option {
	return func(b *Server) {
		b.httpServer = server
	}
}

func WithGRPC(server GRPCServer) Option {
	return func(b *Server) {
		b.grpcServer = server
	}
}

func NewServer(opts ...Option) *Server {
	b := &Server{}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func (b *Server) Run() {
	var wg sync.WaitGroup

	if b.grpcServer != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := b.grpcServer.Run(); err != nil {
				log.Fatalf("gRPC server error: %v", err)
			}
		}()
	}

	if b.httpServer != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := b.httpServer.Run(); err != nil {
				log.Fatalf("HTTP server error: %v", err)
			}
		}()
	}

	if b.httpServer == nil && b.grpcServer == nil {
		log.Println("⚠️ No servers defined. Exiting.")
		return
	}

	wg.Wait()
}
