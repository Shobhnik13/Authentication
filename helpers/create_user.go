package helpers

import (
	"auth/config"
	"auth/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (models.User, error) {
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}

	// get the inserted ID and set it to user.ID
	oid, ok := res.InsertedID.(primitive.ObjectID)

	// if type assertion fails
	if !ok {
		return user, err
	}

	user.ID = oid.Hex()
	// fmt.Println("User created with ID:", user.ID)
	return user, nil
}
