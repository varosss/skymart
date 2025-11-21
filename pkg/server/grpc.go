package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCRegister func(s *grpc.Server)

type GRPCServer struct {
	register GRPCRegister
	port     string
}

func NewGRPCServer(port string, register GRPCRegister) *GRPCServer {
	return &GRPCServer{
		port:     port,
		register: register,
	}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	s.register(grpcServer)

	log.Printf("gRPC server listening on %s", s.port)
	return grpcServer.Serve(lis)
}
