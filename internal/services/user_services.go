package services

import (
	"database/sql"
	"errors"
	"fmt"
	"goserver/internal/database"
	"goserver/internal/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers retrieves all users from PostgreSQL
func GetAllUsers() ([]models.DbUser, error) {
	var users []models.DbUser
	err := database.DB.Select(&users, models.UserQueries.GetAll)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	return users, nil
}

// GetUserByID retrieves a user by ID from PostgreSQL
func GetUserByID(id int) (*models.DbUser, error) {
	var user models.DbUser
	err := database.DB.Get(&user, models.UserQueries.GetByID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*models.DbUser, error) {
	var user models.DbUser
	err := database.DB.Get(&user, models.UserQueries.GetByName, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

// CreateUser creates a new user in PostgreSQL
func CreateUser(user *models.DbUser) error {
	var existingUser models.DbUser
	err := database.DB.Get(&existingUser, models.UserQueries.CheckExists, user.Username, user.Email)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("user already exists: %v", err)
	} else if err == nil {
		return fmt.Errorf("error checking for existing user")
	}

	user.VerifyCode = uuid.New().String()
	user.Approved = false

	if user.Role == "" {
		user.Role = models.USER_ROLES["USER"].Name
	}

	salt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	user.Password = string(salt)
	user.VerifyExpires = time.Now().Add(24 * time.Hour)
	err = database.DB.QueryRowx(models.UserQueries.Insert, user.Username, user.Password, user.Email, user.Role, user.VerifyCode, user.VerifyExpires).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	SendVerificationEmail(user.Email, user.Username, user.VerifyCode)
	return nil
}

// UpdateUser updates an existing user in PostgreSQL
func UpdateUser(id int, user *models.DbUser) error {
	err := database.DB.QueryRowx(models.UserQueries.Update, user.Username, user.Email, user.Role, id).
		Scan(&user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

// UpdateUserPassword updates a user's password
func UpdateUserPassword(id int, passwordHash string) error {
	_, err := database.DB.Exec(models.UserQueries.UpdatePassword, passwordHash, id)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

// DeleteUser deletes a user from PostgreSQL
func DeleteUser(id int) error {
	result, err := database.DB.Exec(models.UserQueries.Delete, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// AuthenticateUser checks user credentials
func AuthenticateUser(username string) (*models.DbUser, error) {
	var user models.DbUser
	err := database.DB.Get(&user, models.UserQueries.Authenticate, username)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func VerifyUserEmail(code string) (*models.DbUser, error) {
	var user models.DbUser
	err := database.DB.Get(&user, models.UserQueries.FindByVerificationCode, code)
	if err != nil {
		return nil, errors.New("invalid verification code")
	}

	// Check if verification code has expired
	if !user.VerifyExpires.IsZero() && time.Now().After(user.VerifyExpires) {
		return nil, errors.New("verification code has expired")
	}

	err = database.DB.QueryRowx(models.UserQueries.ApproveUser, user.ID).Scan(&user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update user verification: %v", err)
	}

	// Update user struct to reflect changes
	user.Approved = true
	user.VerifyCode = ""
	user.VerifyExpires = time.Time{}

	SendWelcomeEmail(user.Email, user.Username) // Changed from user.Name to user.Username
	return &user, nil
}
