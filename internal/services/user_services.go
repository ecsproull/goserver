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

	query := `
        SELECT id, user_name, user_email, user_role, user_approved, created_at, updated_at 
        FROM users 
        ORDER BY created_at DESC
    `

	err := database.DB.Select(&users, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	return users, nil
}

// GetUserByID retrieves a user by ID from PostgreSQL
func GetUserByID(id int) (*models.DbUser, error) {
	var user models.DbUser

	query := `
        SELECT id, user_name, user_email, user_role, created_at, updated_at 
        FROM users 
        WHERE id = $1
    `

	err := database.DB.Get(&user, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*models.DbUser, error) {
	var user models.DbUser

	query := `
        SELECT id, user_name, user_password, user_email, user_role, created_at, updated_at 
        FROM users 
        WHERE user_name = $1
    `

	err := database.DB.Get(&user, query, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

// CreateUser creates a new user in PostgreSQL
func CreateUser(user *models.DbUser) error {
	query_exists := `
        SELECT id FROM users WHERE user_name = $1 OR user_email = $2
    `
	var existingUser models.DbUser
	err := database.DB.Get(&existingUser, query_exists, user.Username, user.Email)
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

	query := `
        INSERT INTO users (user_name, user_password, user_email, user_role, user_verify_code, user_verify_expires) 
        VALUES ($1, $2, $3, $4, $5, $6) 
        RETURNING id, created_at, updated_at
    `

	err = database.DB.QueryRowx(query, user.Username, user.Password, user.Email, user.Role, user.VerifyCode, user.VerifyExpires).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	SendVerificationEmail(user.Email, user.Username, user.VerifyCode)
	return nil
}

// UpdateUser updates an existing user in PostgreSQL
func UpdateUser(id int, user *models.DbUser) error {
	query := `
        UPDATE users 
        SET user_name = $1, user_email = $2, user_role = $3, updated_at = CURRENT_TIMESTAMP
        WHERE id = $4
        RETURNING updated_at
    `

	err := database.DB.QueryRowx(query, user.Username, user.Email, user.Role, id).
		Scan(&user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

// UpdateUserPassword updates a user's password
func UpdateUserPassword(id int, passwordHash string) error {
	query := `
        UPDATE users 
        SET user_password = $1, updated_at = CURRENT_TIMESTAMP
        WHERE id = $2
    `

	_, err := database.DB.Exec(query, passwordHash, id)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

// DeleteUser deletes a user from PostgreSQL
func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := database.DB.Exec(query, id)
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

	query := `
        SELECT id, user_name, user_password, user_email, user_role, created_at, updated_at 
        FROM users 
        WHERE user_name = $1
    `

	err := database.DB.Get(&user, query, username)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func VerifyUserEmail(code string) (*models.DbUser, error) {
	var user models.DbUser

	// First, find the user by verification code
	query := `
        SELECT id, user_name, user_password, user_email, user_role, 
               user_approved, user_verify_code, user_verify_expires
        FROM users 
        WHERE user_verify_code = $1
    `

	err := database.DB.Get(&user, query, code)
	if err != nil {
		return nil, errors.New("invalid verification code")
	}

	// Check if verification code has expired
	if !user.VerifyExpires.IsZero() && time.Now().After(user.VerifyExpires) {
		return nil, errors.New("verification code has expired")
	}

	// Approve user and clear verification code
	updateQuery := `
        UPDATE users 
        SET user_approved = true, 
            user_verify_code = NULL, 
            user_verify_expires = NULL, 
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $1
        RETURNING updated_at
    `

	err = database.DB.QueryRowx(updateQuery, user.ID).Scan(&user.UpdatedAt)
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
