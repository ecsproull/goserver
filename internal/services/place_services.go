package services

import (
	"goserver/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get all places
func GetPlaces() ([]models.Place, error) {
	collection, ctx, cancel := GetCollectionAndContext("places")
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var places []models.Place
	for cursor.Next(ctx) {
		var place models.Place
		if err := cursor.Decode(&place); err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, cursor.Err()
}

// Save a new place
func SavePlace(place *models.Place) error {
	collection, ctx, cancel := GetCollectionAndContext("places")
	defer cancel()

	place.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, place)
	return err
}

// Update a place
func UpdatePlace(id string, updateData map[string]interface{}) (*models.Place, error) {
	collection, ctx, cancel := GetCollectionAndContext("places")
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	delete(updateData, "_id") // Don't allow updating the ID

	update := bson.M{"$set": updateData}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, opts)
	var updated models.Place
	if err := result.Decode(&updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// Delete a place
func DeletePlace(id string) (*models.Place, error) {
	collection, ctx, cancel := GetCollectionAndContext("places")
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var deleted models.Place
	err = collection.FindOneAndDelete(ctx, bson.M{"_id": objID}).Decode(&deleted)
	if err != nil {
		return nil, err
	}
	return &deleted, nil
}
