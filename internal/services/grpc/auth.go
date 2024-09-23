package grpc

import (
	"context"

	models "github.com/ivanbatutin921/Anti-bruteforce/internal/models"

	db "github.com/ivanbatutin921/Anti-bruteforce/internal/database/postgresql"
	pb "github.com/ivanbatutin921/Anti-bruteforce/internal/protobuf"
	service "github.com/ivanbatutin921/Anti-bruteforce/internal/services"
)

type Server struct {
	pb.UnimplementedBruteforceServiceServer
	tokenBuckets map[string]*service.TokenBucket
}

func (s *Server) Authorization(ctx context.Context, req *pb.AuthRequest) (*pb.Response, error) {

	tb := service.NewTokenbucket(3, 10.0) // 1 запрос в 10 секунд
	tbManager := service.NewTokenBucketManager()
	tbManager.AddBucketMemory(req.Login, req.Ip, tb)

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
		return &pb.Response{Ok: false}, nil
	}

	if err := db.CreateUser(&db.PostgreSQLDB{}, &auth); err != nil {
		return &pb.Response{Ok: false}, nil
	}

	return &pb.Response{Ok: true}, nil
}

func (s *Server) ResetBucket(ctx context.Context, req *pb.BucketRequest) (*pb.Response, error) {
	key := req.Login + req.Ip
	tb, ok := s.tokenBuckets[key]
	if !ok {
		return &pb.Response{Ok: false}, nil
	}
	tb.Reset()
	delete(s.tokenBuckets, key)
	return &pb.Response{Ok: true}, nil
}
