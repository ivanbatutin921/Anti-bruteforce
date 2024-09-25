package grpc

import (
	"context"

	db "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/database/postgresql"
	models "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/models"
	pb "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/protobuf"
	service "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/services"
)

type Server struct {
	pb.UnimplementedBruteforceServiceServer
	tbManager *service.TokenBucketManager
}

func NewServer() *Server {
	return &Server{
		tbManager: service.NewTokenBucketManager(),
	}
}

func (s *Server) Authorization(ctx context.Context, req *pb.AuthRequest) (*pb.Response, error) {
	tb := service.NewTokenbucket(3, 10.0) // 1 запрос в 10 секунд
	tbManager := service.NewTokenBucketManager()

	err := tbManager.AddBucketMemory(req, tb)
	if err != nil {
		return &pb.Response{Ok: false}, err
	}

	flag := service.CheckIp(req.Ip)
	if !flag {
		return &pb.Response{Ok: false}, nil
	}

	if !tb.Take(req.Ip, 1) {
		return &pb.Response{Ok: false}, nil
	}

	auth := models.Auth{
		Login:    req.Login,
		Password: req.Password,
		Ip:       req.Ip,
	}

	if err := db.CheckLogin(&db.PostgreSQLDB{}, &auth); err != nil {
		return &pb.Response{Ok: false}, err
	}

	if err := db.CreateUser(&db.PostgreSQLDB{}, &auth); err != nil {
		return &pb.Response{Ok: false}, err
	}

	return &pb.Response{Ok: true}, nil
}

func (s *Server) ResetBucket(ctx context.Context, req *pb.AuthRequest) (*pb.Response, error) {
	bucketReq := &pb.BucketRequest{
		Login: req.Login,
		Ip:    req.Ip,
	}
	err := s.tbManager.ResetBucket(bucketReq)
	if err != nil {
		return &pb.Response{Ok: false}, err
	}
	return &pb.Response{Ok: true}, nil
}
