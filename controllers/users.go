package controllers

import (
	"fmt"

	"github.com/Ericarthurc/GoFiber-Rest-API-Mongo_Template/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
GetUsers | @Desc: Get all users |
@Method: GET |
@Route: "api/v1/users" |
@Auth: Public
*/
func GetUsers(c *fiber.Ctx) error {
	filter := bson.D{{}}
	cursor, err := models.UserCollection.Find(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err})
	}

	var users []models.User = make([]models.User, 0)

	if err := cursor.All(c.Context(), &users); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"success": false, "data": err})
	}

	return c.JSON(fiber.Map{"success": true, "data": users})
}

/*
GetUser | @Desc: Get user by id |
@Method: GET |
@Route: "api/v1/users/:id" |
@Auth: Public
*/
func GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": idParam + " is not a valid id!"})
	}

	filer := bson.D{{Key: "_id", Value: userID}}
	userRecord := models.UserCollection.FindOne(c.Context(), filer)
	if userRecord.Err() != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": "No user with id: " + idParam + " was found!"})
	}

	user := &models.User{}
	userRecord.Decode(user)

	return c.JSON(fiber.Map{"success": true, "data": user})
}

/*
CreateUser | @Desc: Create new user |
@Method: POST |
@Route: "api/v1/users" |
@Auth: Public
*/
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": err})
	}

	insertionResult, err := models.UserCollection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": err})
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := models.UserCollection.FindOne(c.Context(), filter)

	// decode the Mongo record into Employee
	createdUser := &models.User{}
	createdRecord.Decode(createdUser)

	return c.JSON(fiber.Map{"success": true, "data": createdUser})
}

/*
UpdateUser | @Desc: Update user by id |
@Method: PATCH |
@Route: "api/v1/users/:id" |
@Auth: Private
*/
func UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": idParam + " is not a valid id!"})
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": err})
	}

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: user.Name},
				{Key: "age", Value: user.Email},
			},
		},
	}

	after := options.After
	userRecord := models.UserCollection.FindOneAndUpdate(c.Context(), filter, update, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	})

	if userRecord.Err() != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": "No user with id: " + idParam + " was found!"})
	}

	updatedUser := &models.User{}
	userRecord.Decode(updatedUser)

	return c.JSON(fiber.Map{"success": true, "data": updatedUser})
}

/*
DeleteUser | @Desc: Delete user by id |
@Method: DELETE |
@Route: "api/v1/users/:id" |
@Auth: Private
*/
func DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": idParam + " is not a valid id!"})
	}

	filer := bson.D{{Key: "_id", Value: userID}}
	userRecord := models.UserCollection.FindOneAndDelete(c.Context(), filer)
	if userRecord.Err() != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": "No user with id: " + idParam + " was found!"})
	}

	return c.JSON(fiber.Map{"success": true, "data": "User was deleted!"})
}
