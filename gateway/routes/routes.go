package routes

import (
	"context"

	pb "github.com/ivanbatutin921/Anti-bruteforce/internal/protobuf"

	"github.com/gofiber/fiber/v2"
)

type ChatServiceHandler struct {
	pb.BruteforceServiceServer
	mk pb.BruteforceServiceClient
}

func ServiceHandler(mk pb.BruteforceServiceClient) *ChatServiceHandler {
	return &ChatServiceHandler{mk: mk}
}

func (c *ChatServiceHandler) CreateUser(ctx *fiber.Ctx) error {
	var data pb.AuthRequest
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	ctxAuth := context.WithValue(context.Background(),"authBody", data)

	receivedData, err := c.mk.Authorization(ctxAuth, &data)
	if err != nil {
		return err
	}


	return ctx.JSON(receivedData) 
}
