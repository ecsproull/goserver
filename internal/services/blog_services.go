package services

import (
	"context"
	"log"
	"time"

	"goserver/internal/database"
	"goserver/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllBlogs() ([]models.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database("edandlinda").Collection("blogs")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var blogs []models.Blog
	for cursor.Next(ctx) {
		var blog models.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func GetBlogByID(id string) (*models.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database("edandlinda").Collection("blogs")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var blog models.Blog
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&blog)
	if err != nil {
		return nil, nil // Not found or decode error
	}
	return &blog, nil
}

// SaveBlog creates a new blog or updates an existing one based on blog_id
func SaveBlog(data *models.Blog) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database("edandlinda").Collection("blogs")

	if data.ID != primitive.NilObjectID {
		// Update existing blog
		blogID := data.ID
		updateData := *data
		updateData.ID = primitive.NilObjectID // Don't update the ID field

		update := bson.M{
			"$set": updateData,
		}

		result := collection.FindOneAndUpdate(ctx, bson.M{"_id": blogID}, update)

		var updatedBlog models.Blog
		if err := result.Decode(&updatedBlog); err != nil {
			return "", err
		}
		log.Printf("Saved Blog: %s", updatedBlog.Subject)
		return updatedBlog.ID.Hex(), nil
	} else {
		// Create new blog
		log.Printf("Creating new blog post.")
		data.ID = primitive.NewObjectID()
		res, err := collection.InsertOne(ctx, data)
		if err != nil {
			return "", err
		}
		log.Printf("Saved Blog: %s", data.Subject)
		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			return oid.Hex(), nil
		}
		return "", nil
	}
}

// DeleteBlog deletes a blog by its ID
func DeleteBlog(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database("edandlinda").Collection("blogs")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
