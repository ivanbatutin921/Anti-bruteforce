package grpc

import (
	"context"
	"log"

	database "github.com/ivanbatutin921/Anti-bruteforce/mk/service/database"
	models "github.com/ivanbatutin921/Anti-bruteforce/mk/service/models"
	"github.com/ivanbatutin921/Anti-bruteforce/mk/service/pkg/logger"
	pb "github.com/ivanbatutin921/Anti-bruteforce/mk/service/protobuf"
	service "github.com/ivanbatutin921/Anti-bruteforce/mk/service/services"
	"gorm.io/gorm"
)

var db = database.DBDB
var manager = service.NewTokenBucketManager()

type Server struct {
	pb.UnimplementedBruteforceServiceServer
	tbManager *service.TokenBucketManager
	logger    *logger.Logger
	db        *gorm.DB
}

func NewServer(db *gorm.DB, logger *logger.Logger) *Server {
	return &Server{
		tbManager: service.NewTokenBucketManager(),
		logger:    logger,
		db:        db,
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

	flag := s.CheckIp(req.Ip)
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

	existingUser, err := s.CheckLogin(&models.Auth{Login: req.Login})
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
		if err := s.CreateUser(&auth); err != nil {
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
	err := s.CreateWhiteList(&whiteList)
	if err != nil {
		return &pb.WhiteList{Ip: ""}, err
	}
	return &pb.WhiteList{Ip: req.Ip}, nil
}

func (s *Server) DeleteToWhitelist(ctx context.Context, req *pb.WhiteList) (*pb.WhiteList, error) {
	err := s.DeleteWhiteList(req.Ip)
	if err != nil {
		return &pb.WhiteList{Ip: ""}, err
	}
	return &pb.WhiteList{Ip: req.Ip}, nil
}

func (s *Server) AddToBlacklist(ctx context.Context, req *pb.BlackList) (*pb.BlackList, error) {
	blackList := models.BlackList{
		Ip: req.Ip,
	}
	err := s.CreateBlackList(&blackList)
	if err != nil {
		return &pb.BlackList{Ip: ""}, err
	}
	return &pb.BlackList{Ip: req.Ip}, nil
}

func (s *Server) DeleteToBlacklist(ctx context.Context, req *pb.BlackList) (*pb.BlackList, error) {
	err := s.DeleteBlackList(req.Ip)
	if err != nil {
		return &pb.BlackList{Ip: ""}, err
	}
	return &pb.BlackList{Ip: req.Ip}, nil
}
