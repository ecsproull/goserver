package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"goserver/internal/database"
	"goserver/internal/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func getUsersCollectionAndContext() (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoClient.Database("edandlinda").Collection("users")
	return collection, ctx, cancel
}

func GetUserByID(id string) (*models.User, error) {
	collection, ctx, cancel := getUsersCollectionAndContext()
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	collection, ctx, cancel := getUsersCollectionAndContext()
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database("edandlinda").Collection("users")

	// Check if email already exists
	if user.Email == "" {
		return errors.New("email is required")
	}

	var existingUser models.User
	err := collection.FindOne(ctx, bson.M{"user_email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("email already in use")
	} else if err != mongo.ErrNoDocuments {
		return fmt.Errorf("error checking email: %v", err)
	}

	// Check if username already exists
	err = collection.FindOne(ctx, bson.M{"user_name": user.Name}).Decode(&existingUser)
	if err == nil {
		return errors.New("username already in use")
	} else if err != mongo.ErrNoDocuments {
		return fmt.Errorf("error checking username: %v", err)
	}

	// Generate verification code
	user.VerifyCode = uuid.New().String()
	user.Approved = false // Default to false until verified

	// Set default role if not provided
	if user.Role == "" {
		user.Role = models.USER_ROLES["USER"].Name
	}

	// Set verify expiration (24 hours from now)
	user.VerifyExpires = time.Now().Add(24 * time.Hour)

	// Hash password
	salt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	user.Password = string(salt)

	// Set timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Insert user
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		// Check for MongoDB duplicate key error
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("duplicate key error")
		}
		return fmt.Errorf("error saving user: %v", err)
	}

	// Set the ID from the result
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	go func(u *models.User) {
		_ = SendVerificationEmail(u.Email, u.Name, u.VerifyCode)
	}(user)

	fmt.Printf("Saved User: %s\n", user.Name)
	return nil
}

func UpdateUser(id string, user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database("edandlinda").Collection("users")

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID: %v", err)
	}

	// Build update document
	updateDoc := bson.M{}

	// If password is being updated, hash it
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error hashing password: %v", err)
		}
		updateDoc["user_password"] = string(hashedPassword)
	}

	// Add other fields to update (exclude empty values)
	if user.Name != "" {
		updateDoc["user_name"] = user.Name
	}
	if user.Email != "" {
		updateDoc["user_email"] = user.Email
	}
	if user.Role != "" {
		updateDoc["role"] = user.Role
	}
	// Add Approved even if false (it's a boolean)
	updateDoc["user_approved"] = user.Approved

	// Set updated timestamp
	updateDoc["updatedAt"] = time.Now()

	// Perform update
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updateDoc},
	)

	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	fmt.Printf("Updated User: %s\n", user.Name)
	return nil
}

func DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database("edandlinda").Collection("users")

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID: %v", err)
	}

	// Delete the user
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	fmt.Printf("Deleted user with ID: %s\n", id)
	return nil
}

func VerifyUserEmail(code string) (*models.User, error) {
	collection, ctx, cancel := GetCollectionAndContext("users")
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"user_verify_code": code}).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid or expired verification code")
	}

	// Check if verification code has expired
	if !user.VerifyExpires.IsZero() && time.Now().After(user.VerifyExpires) {
		return nil, errors.New("verification code has expired")
	}

	// Approve user and clear verification code
	update := bson.M{
		"$set": bson.M{
			"approved":       true,
			"verify_code":    nil,
			"verify_expires": nil,
		},
	}
	_, err = collection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		return nil, err
	}

	// Update user struct to reflect changes
	user.Approved = true
	user.VerifyCode = ""
	user.VerifyExpires = time.Time{}

	SendWelcomeEmail(user.Email, user.Name)
	return &user, nil
}
