package grpc

import (
	"context"

	db "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/database/postgresql"
	models "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/models"
	service "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/services"
	pb "github.com/ivanbatutin921/Anti-bruteforce/protobuf"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBruteforceServiceServer
	tbManager *service.TokenBucketManager
}

// RegisterService implements grpc.ServiceRegistrar.
func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl any) {
	panic("unimplemented")
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

func (s *Server) ResetBucket(ctx context.Context, req *pb.BucketRequest) (*pb.Response, error) {

	err := s.tbManager.ResetBucket(req)
	if err != nil {
		return &pb.Response{Ok: false}, err
	}
	return &pb.Response{Ok: true}, nil
}

func (s *Server) AddToBlacklist(ctx context.Context, req *pb.BlackList) (*pb.BlackList, error) {
	blackList := models.BlackList{
		Ip: req.Ip,
	}
	err := db.CreateBlackList(&db.PostgreSQLDB{}, &blackList)
	if err != nil {
		return &pb.BlackList{Ip: ""}, err
	}
	return &pb.BlackList{Ip: req.Ip}, nil
}

func (s *Server) DeleteToBlacklist(ctx context.Context, req *pb.BlackList) (*pb.BlackList, error) {
	err := db.DeleteBlackList(&db.PostgreSQLDB{}, req.Ip)
	if err != nil {
		return &pb.BlackList{Ip: ""}, err
	}
	return &pb.BlackList{Ip: req.Ip}, nil
}

func (s *Server) AddToWhitelist(ctx context.Context, req *pb.WhiteList) (*pb.WhiteList, error) {
	whiteList := models.WhiteList{
		Ip: req.Ip,
	}
	err := db.CreateWhiteList(&db.PostgreSQLDB{}, &whiteList)
	if err != nil {
		return &pb.WhiteList{Ip: ""}, err
	}
	return &pb.WhiteList{Ip: req.Ip}, nil
}

func (s *Server) DeleteToWhitelist(ctx context.Context, req *pb.WhiteList) (*pb.WhiteList, error) {
	err := db.DeleteWhiteList(&db.PostgreSQLDB{}, req.Ip)
	if err != nil {
		return &pb.WhiteList{Ip: ""}, err
	}
	return &pb.WhiteList{Ip: req.Ip}, nil
}
