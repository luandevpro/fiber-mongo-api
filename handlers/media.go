package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	_, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("upload", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	file, err := c.FormFile("document")
	if err != nil {
		return err
	}
	// Save file to root directory:
	return c.SaveFile(file, fmt.Sprintf("./upload/%s", file.Filename))
}

func ViewMedia(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendFile("./upload/" + id)
}

func DeleteMedia(c *fiber.Ctx) error {
	id := c.Params("id")
	path := "upload/" + id

	err := os.Remove(path)

	if err != nil {
		return c.Status(404).SendString("File not found")
	}

	return c.Status(200).JSON(fiber.Map{"msg": "file deleted"})
}
