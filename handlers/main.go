package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

type Fullname struct {
	Name string
	Age  int
}

func GetHome(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}
