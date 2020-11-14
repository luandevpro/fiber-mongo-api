package routers

import (
	"github.com/gofiber/fiber/v2"

	"fibermongo/handlers"
	"fibermongo/middlewares"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.GetHome)
	// login
	app.Post("/login", handlers.Login)
	// api user
	app.Get("/user", handlers.GetAllUser)
	app.Get("/user/:id", handlers.GetUser)
	app.Post("/user", handlers.CreateUser)
	app.Put("/user/:id", handlers.UpdateUser)
	app.Delete("/user/:id", handlers.DeleteUser)
	// api profile
	app.Get("/profile", middlewares.AuthRequired(), handlers.Profile)
	// api user
	app.Get("/post", handlers.GetAllPost)
	app.Get("/post/:id", handlers.GetPost)
	app.Post("/post", handlers.CreatePost)
	app.Put("/post/:id", handlers.UpdatePost)
	app.Delete("/post/:id", handlers.DeletePost)
}
