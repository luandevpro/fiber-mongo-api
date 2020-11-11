package main

import (
	"github.com/gofiber/fiber/v2"

	"fibermongo/databases"
	"fibermongo/routers"
)

func main() {

	databases.InitDatabase()

	app := fiber.New()

	routers.Setup(app)

	app.Listen(":3000")

}
