package grpc

import (
	"log"
	"net"
	pb "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/protobuf"
)

func ListenGRPC() {
	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := NewServer()
	pb.RegisterBruteforceServiceServer(s, &Server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
