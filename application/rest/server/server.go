package server

import "github.com/gofiber/fiber/v2"

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{})
	return app
}
