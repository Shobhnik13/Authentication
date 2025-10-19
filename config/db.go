package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() error {
	// its just like require("dotenv").config() in nodejs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	// out of full env accessing the mongo uri and db name
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	if uri == "" || dbName == "" {
		log.Fatal("MONGO_URI or DB_NAME not set in .env file")
		return err
	}

	// setting the context window for 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connecting now
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Mongo connection error:", err)
		return err
	}

	// pinging to check
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo ping error:", err)
		return err
	}

	// assigning the database instance to the global variable
	DB = client.Database(dbName)
	log.Println("Connected to MongoDB!")
	return nil
}
