package grpc

import (
	"log"
	"net"

	pb "github.com/ivanbatutin921/Anti-bruteforce/protobuf"
	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {
	s := NewServer()
	grpcServer := grpc.NewServer()
	pb.RegisterBruteforceServiceServer(grpcServer, s)

	return grpcServer
}

func ListenGRPC() {
	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := NewGRPCServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
