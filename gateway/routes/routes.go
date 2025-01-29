package routes

import (
	"context"

	pb "github.com/ivanbatutin921/Anti-bruteforce/protobuf"

	"github.com/gofiber/fiber/v2"
)

type ChatServiceHandler struct {
	pb.BruteforceServiceServer
	mk pb.BruteforceServiceClient
}

func ServiceHandler(mk pb.BruteforceServiceClient) *ChatServiceHandler {
	return &ChatServiceHandler{mk: mk}
}

func (c *ChatServiceHandler) Authorization(ctx *fiber.Ctx) error {
	var data pb.AuthRequest
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	receivedData, err := c.mk.Authorization(context.Background(), &data)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if receivedData == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(fiber.Map{"ok": receivedData.Ok})
}

func (c *ChatServiceHandler) ResetBucket(ctx *fiber.Ctx) error {
	var data pb.BucketRequest
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	receivedData, err := c.mk.ResetBucket(context.Background(), &data)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if receivedData == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(fiber.Map{"ok": receivedData.Ok})
}

func (c *ChatServiceHandler) AddToBlacklist(ctx *fiber.Ctx) error {
	var data pb.BlackList
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	receivedData, err := c.mk.AddToBlacklist(context.Background(), &data)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if receivedData == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(fiber.Map{"ip": receivedData.Ip})
}

func (c *ChatServiceHandler) DeleteToBlacklist(ctx *fiber.Ctx) error {
	var data pb.BlackList
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	receivedData, err := c.mk.DeleteToBlacklist(context.Background(), &data)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if receivedData == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(fiber.Map{"ip": receivedData.Ip})
}

func (c *ChatServiceHandler) AddToWhitelist(ctx *fiber.Ctx) error {
	var data pb.WhiteList
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	receivedData, err := c.mk.AddToWhitelist(context.Background(), &data)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if receivedData == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(fiber.Map{"ip": receivedData.Ip})
}

func (c *ChatServiceHandler) DeleteToWhitelist(ctx *fiber.Ctx) error {
	var data pb.WhiteList
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	receivedData, err := c.mk.DeleteToWhitelist(context.Background(), &data)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if receivedData == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(fiber.Map{"ip": receivedData.Ip})
}
