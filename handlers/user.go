package handlers

import (
	"fibermongo/databases"
	"fibermongo/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(c *fiber.Ctx) error {
	p := new(models.User)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	collection := databases.Db.Collection("user")

	// insert one user to collections
	insertOne, _ := collection.InsertOne(c.Context(), p)

	// id := insertOne.InsertedID.(primitive.ObjectID).Hex()

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertOne.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into User
	createdUser := &models.User{}
	createdRecord.Decode(createdUser)

	return c.Status(201).JSON(createdUser)
}

func GetAllUser(c *fiber.Ctx) error {
	collection := databases.Db.Collection("user")

	cursor, err := collection.Find(c.Context(), bson.D{{}}, options.Find().SetLimit(2))

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []models.User = make([]models.User, 0)

	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(idParam)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(404)
	}

	// Find the employee and update its data
	filter := bson.D{{Key: "_id", Value: userID}}

	collection := databases.Db.Collection("user")

	var resultOne models.User

	// find one documents
	err = collection.FindOne(c.Context(), filter).Decode(&resultOne)

	if err != nil {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(resultOne)
}

func UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(idParam)

	p := new(models.User)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(404)
	}

	// Find the employee and update its data
	filter := bson.D{{Key: "_id", Value: userID}}

	collection := databases.Db.Collection("user")

	var resultOne models.User

	// find one documents
	err = collection.FindOne(c.Context(), filter).Decode(&resultOne)

	if err != nil {
		return c.SendStatus(404)
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "email", Value: p.Email},
			{Key: "age", Value: p.Age},
			{Key: "password", Value: p.Password},
			{Key: "name", Value: p.Name},
			{Key: "status", Value: p.Status},
		}},
	}

	// update one documents with filter and data update
	err = collection.FindOneAndUpdate(c.Context(), filter, update).Err()

	if err != nil {
		return c.SendStatus(400)
	}

	p.ID = userID

	return c.Status(200).JSON(p)
}

func DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userId, err := primitive.ObjectIDFromHex(idParam)

	filter := bson.D{{Key: "_id", Value: userId}}

	collection := databases.Db.Collection("user")
	
	result, err := collection.DeleteOne(c.Context(), &filter)
	
	if err != nil {
		return c.SendStatus(500)
	}

	// the employee might not exist
	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	// the record was deleted
	return c.SendStatus(204)

}
