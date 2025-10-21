package helpers

import (
	"auth/config"
	"auth/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func GetUserByEmail(email string) (models.User, error) {
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	// create filter for search
	filter := bson.M{"email": email}

	res := collection.FindOne(ctx, filter)

	// now decoding the result into user
	err := res.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(id string) (models.User, error) {
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	// convert string id to ObjectID for matching on mongodb indexed field of _id
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, errors.New("invalid user ID")
	}

	// create filter for search
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)

	// now decoding the result into user
	err = res.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
