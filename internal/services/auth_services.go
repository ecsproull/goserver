package services

import (
	"errors"
	"goserver/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateLoginInput validates the login input
func ValidateLoginInput(userName, userPassword string) []ValidationError {
	var errs []ValidationError

	if userName == "" {
		errs = append(errs, ValidationError{Field: "user_name", Message: "Username is required"})
	}

	if userPassword == "" {
		errs = append(errs, ValidationError{Field: "user_password", Message: "Password is required"})
	}

	return errs
}

// LoginUser handles user authentication and approval check
func LoginUser(userName, userPassword string) (*models.User, []ValidationError, error) {
	// Validate input
	validationErrors := ValidateLoginInput(userName, userPassword)

	if len(validationErrors) > 0 {
		return nil, validationErrors, nil
	}

	// Authenticate user (replace with your DB lookup)
	foundUser, err := GetUser(userName, userPassword)
	if err != nil {
		return nil, nil, err
	}

	// Check if user is approved (email verified)
	if !foundUser.UserApproved {
		return nil, nil, errors.New("please verify your email address before logging in")
	}

	return foundUser, nil, nil
}

// Dummy GetUser for illustration; replace with real DB logic
func GetUser(userName, userPassword string) (*models.User, error) {
	collection, ctx, cancel := GetCollectionAndContext("users")
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{
		"user_name": userName,
	}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil // User not found
	}
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the hashed password in the database
	if bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(userPassword)) != nil {
		return nil, nil // Password does not match
	}

	return &user, nil
}
