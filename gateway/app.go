package app

import (
	"context"
	"log"
	"time"

	routes "github.com/ivanbatutin921/Anti-bruteforce/gateway/routes"
	pb "github.com/ivanbatutin921/Anti-bruteforce/protobuf"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	app  *fiber.App
	mk   pb.BruteforceServiceClient
	port string
}

func (s *Server) runGrpcServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не получилось соединиться: %v", err)
	}
	s.mk = pb.NewBruteforceServiceClient(conn)
}

func (s *Server) allRoutes() {
	bruteforce := routes.ServiceHandler(s.mk)

	s.app.Post("/auth", bruteforce.Authorization)

	s.app.Post("/bucket/reset", bruteforce.ResetBucket)

	s.app.Post("/whitelist/add", bruteforce.AddToWhitelist)

	s.app.Delete("/whitelist/delete", bruteforce.DeleteToWhitelist)

	s.app.Post("/blacklist/add", bruteforce.AddToBlacklist)

	s.app.Delete("/blacklist/delete", bruteforce.DeleteToBlacklist)

	s.app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
func NewServer(port string) *Server {
	s := &Server{
		app:  fiber.New(),
		port: port,
	}
	s.app.Use(logger.New())
	return s
}

func (s *Server) Run() {
	s.runGrpcServer()
	s.allRoutes()
	log.Fatal(s.app.Listen(":" + s.port))
}

func App() {
	s := NewServer("3000")
	s.Run()
}
