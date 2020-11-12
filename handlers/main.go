package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}
