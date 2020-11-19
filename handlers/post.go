package handlers

import (
	"fibermongo/databases"
	"fibermongo/models"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/satori/go.uuid"
	"log"
)

func CreatePost(c *fiber.Ctx) error {
	p := new(models.Post)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	databases.Db.Create(&p)

	return c.Status(201).JSON(p)
}

func GetAllPost(c *fiber.Ctx) error {
	var posts []models.Post

	databases.Db.Preload("User").Find(&posts)

	return c.Status(200).JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	postId := c.Params("id")

	var post models.Post

	id , _ := uuid.FromString(postId)

	databases.Db.Where("id = ?", id).Preload("User").Find(&post)

	return c.Status(200).JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	postId := c.Params("id")
	log.Println(postId)

	p := new(models.Post)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	var post models.Post

	id, _ := uuid.FromString(postId)

	p.ID = id

	log.Println("id", id)

	databases.Db.Where("id = ?", id).First(&post).Update(&p)

	return c.Status(200).JSON(p)
}

func DeletePost(c *fiber.Ctx) error {
	postId := c.Params("id")

	id, _ := uuid.FromString(postId)

	var post models.Post

	databases.Db.Where("id = ?", id).First(&post).Delete(&post)

	// the record was deleted
	return c.SendStatus(204)

}
