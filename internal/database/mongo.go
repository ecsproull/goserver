package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient is the exported client instance
var MongoClient *mongo.Client

// InitMongo initializes the MongoDB client and returns an error if it fails.
// Pass the MongoDB URI as the argument.
func InitMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return err
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		return err
	}

	MongoClient = client
	fmt.Println("Connected to MongoDB successfully.")
	return nil
}
