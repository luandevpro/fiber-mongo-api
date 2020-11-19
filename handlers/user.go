package handlers

import (
	"fibermongo/databases"
	"fibermongo/models"
	"fibermongo/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	p := new(models.User)

	argon2ID := utils.NewArgon2ID()

	if err := c.BodyParser(p); err != nil {
		return err
	}

	hash, _ := argon2ID.Hash(p.Password)

	p.Password = hash

	databases.Db.Create(&p)

	return c.Status(201).JSON(p)
}

func GetAllUser(c *fiber.Ctx) error {
	var users []models.User

	databases.Db.Find(&users)

	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user models.User

	databases.Db.First(&user, userId)

	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var user models.User

	p := new(models.User)

	argon2ID := utils.NewArgon2ID()

	if err := c.BodyParser(p); err != nil {
		return err
	}

	hash, _ := argon2ID.Hash(p.Password)

	id, _ := strconv.ParseUint(userID, 10, 64)

	p.ID = id

	p.Password = hash

	databases.Db.First(&user, userID).Update(&p)

	return c.Status(200).JSON(p)
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user models.User

	databases.Db.First(&user, userId)

	if user.Email == "" {
		return c.Status(500).SendString("No Book Found with ID")
	}

	databases.Db.Delete(&user)

	return c.SendStatus(204)

}

func Profile(c *fiber.Ctx) error {
	return c.SendString("Welcome ")
}
