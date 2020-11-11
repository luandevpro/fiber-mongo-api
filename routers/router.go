package routers

import (
	"github.com/gofiber/fiber/v2"

	"fibermongo/handlers"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.GetHome)
	// api user
	app.Get("/user", handlers.GetAllUser)
	app.Get("/user/:id", handlers.GetUser)
	app.Post("/user", handlers.CreateUser)
	app.Put("/user/:id", handlers.UpdateUser)
	app.Delete("/user/:id", handlers.DeleteUser)
}
