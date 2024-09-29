package grpc

import (
	"context"
	"log"

	database "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/database/postgresql"
	models "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/models"
	service "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/services"
	pb "github.com/ivanbatutin921/Anti-bruteforce/protobuf"
)

var db = database.DBDB
var manager = service.NewTokenBucketManager()

type Server struct {
	pb.UnimplementedBruteforceServiceServer
	tbManager *service.TokenBucketManager
}

// RegisterService implements grpc.ServiceRegistrar.
// func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl any) {
// 	panic("unimplemented")
// }

func NewServer() *Server {
	return &Server{
		tbManager: service.NewTokenBucketManager(),
	}
}

func (s *Server) Authorization(ctx context.Context, req *pb.AuthRequest) (*pb.Response, error) {
	tbManager := manager

	// Retrieve the existing TokenBucket instance from the manager
	tb, err := tbManager.GetBucket(req.Login, req.Ip)
	log.Printf("tb: %+v\n", tb)
	if err != nil {
		// If the bucket doesn't exist, create a new one
		tb = service.NewTokenbucket(3, 0.1)
		err = tbManager.AddBucketMemory(req, tb)
		if err != nil {
			log.Println(err.Error())
			return &pb.Response{Ok: false}, err
		}
	}

	flag := db.CheckIp(req.Ip)
	if !flag {
		return &pb.Response{Ok: false}, nil
	}

	if !tb.Take() {
		log.Println("Too many requests")
		return &pb.Response{Ok: false}, nil
	}

	auth := models.Auth{
		Login:    req.Login,
		Password: req.Password,
		Ip:       req.Ip,
	}

	existingUser, err := db.CheckLogin(&models.Auth{Login: req.Login})
	if err != nil {
		log.Println(err.Error())
		return &pb.Response{Ok: false}, err
	}

	if existingUser != nil {
		// User exists, check password
		if existingUser.Password == req.Password {
			return &pb.Response{Ok: true}, nil
		} else {
			return &pb.Response{Ok: false}, nil
		}
	} else {
		// User does not exist, create new user
		if err := db.CreateUser(&auth); err != nil {
			log.Println(err.Error())
			return &pb.Response{Ok: false}, err
		}
		return &pb.Response{Ok: true}, nil
	}
}

func (s *Server) ResetBucket(ctx context.Context, req *pb.BucketRequest) (*pb.Response, error) {

	err := s.tbManager.ResetBucket(req)
	if err != nil {
		return &pb.Response{Ok: false}, err
	}
	return &pb.Response{Ok: true}, nil
}

func (s *Server) AddToWhitelist(ctx context.Context, req *pb.WhiteList) (*pb.WhiteList, error) {
	whiteList := models.WhiteList{
		Ip: req.Ip,
	}
	err := db.CreateWhiteList(&whiteList)
	if err != nil {
		return &pb.WhiteList{Ip: ""}, err
	}
	return &pb.WhiteList{Ip: req.Ip}, nil
}

func (s *Server) DeleteToWhitelist(ctx context.Context, req *pb.WhiteList) (*pb.WhiteList, error) {
	err := db.DeleteWhiteList(req.Ip)
	if err != nil {
		return &pb.WhiteList{Ip: ""}, err
	}
	return &pb.WhiteList{Ip: req.Ip}, nil
}

func (s *Server) AddToBlacklist(ctx context.Context, req *pb.BlackList) (*pb.BlackList, error) {
	blackList := models.BlackList{
		Ip: req.Ip,
	}
	err := db.CreateBlackList(&blackList)
	if err != nil {
		return &pb.BlackList{Ip: ""}, err
	}
	return &pb.BlackList{Ip: req.Ip}, nil
}

func (s *Server) DeleteToBlacklist(ctx context.Context, req *pb.BlackList) (*pb.BlackList, error) {
	err := db.DeleteBlackList(req.Ip)
	if err != nil {
		return &pb.BlackList{Ip: ""}, err
	}
	return &pb.BlackList{Ip: req.Ip}, nil
}
