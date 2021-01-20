package models

import (
	"github.com/Ericarthurc/GoFiber-Rest-API-Mongo_Template/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserCollection | @desc: the user ccollection on the database
var UserCollection *mongo.Collection

// User | @desc: user model struct
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Email string             `bson:"email,omitempty"`
}

// CreateUserSchema | @desc: adds schema validation and indexes to collection
func CreateUserSchema() {
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"name", "email"},
		"properties": bson.M{
			"name": bson.M{
				"bsonType":    "string",
				"description": "must be a string and is required",
			},
			"email": bson.M{
				"bsonType":    "string",
				"description": "must be a string and is required",
			},
		},
	}

	validator := bson.M{
		"$jsonSchema": jsonSchema,
	}

	database.DB.CreateCollection(database.Ctx, "users", options.CreateCollection().SetValidator(validator))

	UserCollection = database.DB.Collection("users")

	// Make indexes for user collection
	_, _ = UserCollection.Indexes().CreateOne(database.Ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

}
