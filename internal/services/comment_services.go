package services

import (
	"time"

	"goserver/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCommentsByBlogID(blogID string) ([]models.Comment, error) {
	collection, ctx, cancel := GetCollectionAndContext("comments")
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"blog_id": objID}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

// AddComment adds a new comment to the comments collection
func AddComment(comment *models.Comment) (primitive.ObjectID, error) {
	collection, ctx, cancel := GetCollectionAndContext("comments")
	defer cancel()

	// Set the comment ID and CreatedAt
	comment.ID = primitive.NewObjectID()
	comment.CreatedAt = time.Now()

	res, err := collection.InsertOne(ctx, comment)
	if err != nil {
		return primitive.NilObjectID, err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}
	return primitive.NilObjectID, nil
}

// UpdateComment updates a comment by its ID and blog ID
func UpdateComment(blogID, commentID string, updateData map[string]interface{}) error {
	collection, ctx, cancel := GetCollectionAndContext("comments")
	defer cancel()

	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	commentObjID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}

	// Remove fields that should not be updated
	delete(updateData, "_id")
	delete(updateData, "blog_id")

	filter := bson.M{"_id": commentObjID, "blog_id": blogObjID}
	update := bson.M{"$set": updateData}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

// DeleteComment deletes a comment by its ID and blog ID
func DeleteComment(blogID, commentID string) error {
	collection, ctx, cancel := GetCollectionAndContext("comments")
	defer cancel()

	blogObjID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return err
	}
	commentObjID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": commentObjID, "blog_id": blogObjID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
