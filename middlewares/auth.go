package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

var JwtSecret = "asecret"

// Protected protect routes
func AuthRequired() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(JwtSecret),
	})
}
