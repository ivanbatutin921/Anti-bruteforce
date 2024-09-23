package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ivanbatutin921/Anti-bruteforce/internal/services/protobuf"
)

type GRPCServer struct {
	server *grpc.Server
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
	}
}

func (s *GRPCServer) Start() error {
	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	pb.RegisterBruteforceServiceServer(s.server, &Server{})
	if err := s.server.Serve(lsn); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}

	return nil
}

func (s *GRPCServer) Stop() error {
	s.server.Stop()
	return nil
}
