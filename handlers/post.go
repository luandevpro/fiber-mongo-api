package handlers

import (
	"fibermongo/databases"
	"fibermongo/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(c *fiber.Ctx) error {
	p := new(models.Post)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	collection := databases.Db.Collection("post")

	// insert one post to collections
	insertOne, _ := collection.InsertOne(c.Context(), p)

	// id := insertOne.InsertedID.(primitive.ObjectID).Hex()

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertOne.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into post
	createdPost := &models.Post{}
	createdRecord.Decode(createdPost)

	return c.Status(201).JSON(createdPost)
}

func GetAllPost(c *fiber.Ctx) error {
	collection := databases.Db.Collection("post")

	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "user"}, {Key: "localField", Value: "user"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "user"}}}}
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}

	cursor, err := collection.Aggregate(c.Context(), mongo.Pipeline{lookupStage, unwindStage})

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var posts []bson.M

	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(c.Context(), &posts); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	postId, err := primitive.ObjectIDFromHex(idParam)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(404)
	}

	collection := databases.Db.Collection("post")

	// find one documents
	// err = collection.FindOne(c.Context(), filter).Decode(&resultOne)

	matchStage := bson.D{{
		Key:   "$match",
		Value: bson.D{{Key: "_id", Value: postId}},
	}}

	lookupStage := bson.D{
		{
			Key:   "$lookup",
			Value: bson.D{{Key: "from", Value: "user"}, {Key: "localField", Value: "user"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "user"}},
		},
	}

	unwindStage := bson.D{{
		Key:   "$unwind",
		Value: bson.D{{Key: "path", Value: "$user"}, {Key: "preserveNullAndEmptyArrays", Value: true}},
	}}

	cursor, err := collection.Aggregate(c.Context(), mongo.Pipeline{matchStage, lookupStage, unwindStage})

	var posts []bson.M

	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(c.Context(), &posts); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(posts[0])
}

func UpdatePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	postId, err := primitive.ObjectIDFromHex(idParam)

	p := new(models.Post)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(404)
	}

	// Find the employee and update its data
	filter := bson.D{{Key: "_id", Value: postId}}

	collection := databases.Db.Collection("post")

	var resultOne models.Post

	// find one documents
	err = collection.FindOne(c.Context(), filter).Decode(&resultOne)

	if err != nil {
		return c.SendStatus(404)
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: p.Title},
			{Key: "description", Value: p.Description},
		}},
	}

	// update one documents with filter and data update
	err = collection.FindOneAndUpdate(c.Context(), filter, update).Err()

	if err != nil {
		return c.SendStatus(400)
	}

	p.ID = postId

	return c.Status(200).JSON(p)
}

func DeletePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	postId, err := primitive.ObjectIDFromHex(idParam)

	filter := bson.D{{Key: "_id", Value: postId}}

	collection := databases.Db.Collection("post")

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
