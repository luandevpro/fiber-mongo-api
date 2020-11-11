package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	return c.SendString("hello world")
}
