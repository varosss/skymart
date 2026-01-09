package grpcserver

import (
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	lis    net.Listener
}

func NewServer(addr string, register func(*grpc.Server)) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	grpcSrv := grpc.NewServer()
	register(grpcSrv)

	return &Server{
		server: grpcSrv,
		lis:    lis,
	}, nil
}

func (s *Server) Run() error {
	return s.server.Serve(s.lis)
}

func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}
