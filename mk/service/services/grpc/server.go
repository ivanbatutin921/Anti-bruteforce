package grpc

import (
	"log"
	"net"

	pb "github.com/ivanbatutin921/Anti-bruteforce/protobuf"
	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	//log.Println(grpcServer,&Server{})
	pb.RegisterBruteforceServiceServer(grpcServer, &Server{})

	return grpcServer
}

func ListenGRPC() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := NewGRPCServer()
	if grpcServer == nil {
		log.Fatalf("failed to create gRPC server")
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}

	log.Println("gRPC server is running on", lis.Addr().String())
}