package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"fibermongo/databases"
	"fibermongo/routers"
)

func main() {

	databases.InitDatabase()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routers.Setup(app)

	app.Listen(":8080")

}
