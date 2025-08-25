package services

import (
	"errors"
	"goserver/internal/database"
	"goserver/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
func LoginUser(userName, userPassword string) (*models.DbUser, []ValidationError, error) {
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

	if foundUser == nil {
		return nil, nil, errors.New("invalid username or password")
	}

	// Check if user is approved (email verified)
	if !foundUser.Approved {
		return nil, nil, errors.New("please verify your email address before logging in")
	}

	return foundUser, nil, nil
}

// GetUser authenticates user with PostgreSQL database
func GetUser(userName, userPassword string) (*models.DbUser, error) {
	var user models.DbUser

	query := `
        SELECT id, user_name, user_password, user_email, user_role, user_approved
        FROM users 
        WHERE user_name = $1
    `

	err := database.DB.Get(&user, query, userName)
	if err != nil {
		return nil, nil // User not found or database error
	}

	// Compare the provided password with the hashed password in the database
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPassword)) != nil {
		return nil, nil // Password does not match
	}

	return &user, nil
}

// GenerateAccessToken creates a JWT token for the authenticated user
func GenerateAccessToken(user *models.DbUser) (string, error) {
	claims := jwt.MapClaims{
		"user_name": user.Username,
		"user":      user.ID,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
