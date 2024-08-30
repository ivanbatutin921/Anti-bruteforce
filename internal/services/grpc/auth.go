package grpc

import (
	pb "github.com/ivanbatutin921/Anti-bruteforce/internal/transport/grpc/bruteforce.proto"
	"github.com/ivanbatutin921/Anti-bruteforce/internal/models"
	"github.com/ivanbatutin921/Anti-bruteforce/internal/services"
)

type Server struct{
	pb.U
}