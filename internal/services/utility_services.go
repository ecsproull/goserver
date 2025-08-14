package services

import (
	"context"
	"time"

	"goserver/internal/database"

	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollectionAndContext returns a MongoDB collection, context, and cancel function for a given collection name.
func GetCollectionAndContext(collectionName string) (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoClient.Database("edandlinda").Collection(collectionName)
	return collection, ctx, cancel
}
