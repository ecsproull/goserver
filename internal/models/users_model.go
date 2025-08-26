package models

import (
	"time"
)

type UserRole struct {
	Name  string
	Level int
}

var USER_ROLES = map[string]UserRole{
	"USER":      {Name: "User", Level: 1},
	"MANUALS":   {Name: "Manuals", Level: 2},
	"COMMENTOR": {Name: "Commentor", Level: 3},
	"CREATOR":   {Name: "Creator", Level: 4},
	"ADMIN":     {Name: "Admin", Level: 5},
}

// PostgreSQL structures
type DbUser struct {
	ID            int       `json:"id" db:"id"`
	Username      string    `json:"user_name" db:"user_name"`
	Password      string    `json:"user_password" db:"user_password"`
	Email         string    `json:"user_email" db:"user_email"`
	Role          string    `json:"user_role" db:"user_role"`
	VerifyCode    string    `json:"user_verify_code" db:"user_verify_code"`
	VerifyExpires time.Time `json:"user_verify_expires" db:"user_verify_expires"`
	Approved      bool      `json:"user_approved" db:"user_approved"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type UQueries struct {
	GetAll                 string
	GetByID                string
	GetByName              string
	CheckExists            string
	Insert                 string
	Update                 string
	UpdatePassword         string
	Delete                 string
	Authenticate           string
	FindByVerificationCode string
	ApproveUser            string
}

var UserQueries = UQueries{
	GetAll: `
			SELECT id, user_name, user_email, user_role, user_approved, created_at, updated_at 
      FROM users 
      ORDER BY created_at DESC
    `,
	GetByID: `
 			SELECT id, user_name, user_email, user_role, created_at, updated_at 
      FROM users 
      WHERE id = $1
		`,
	GetByName: `
			SELECT id, user_name, user_password, user_email, user_role, created_at, updated_at 
      FROM users 
      WHERE user_name = $1
		`,
	CheckExists: `
			SELECT id FROM users WHERE user_name = $1 OR user_email = $2
		`,
	Insert: `
			INSERT INTO users (user_name, user_password, user_email, user_role, user_verify_code, user_verify_expires) 
      VALUES ($1, $2, $3, $4, $5, $6) 
      RETURNING id, created_at, updated_at
    `,
	Update: `
			UPDATE users 
      SET user_name = $1, user_email = $2, user_role = $3, updated_at = CURRENT_TIMESTAMP
      WHERE id = $4
      RETURNING updated_at
    `,
	UpdatePassword: `
			UPDATE users 
      SET user_password = $1, updated_at = CURRENT_TIMESTAMP
      WHERE id = $2
	`,
	Delete: `
			DELETE FROM users WHERE id = $1
    `,
	Authenticate: `
      SELECT id, user_name, user_password, user_email, user_role, created_at, updated_at 
      FROM users 
      WHERE user_name = $1
	`,
	FindByVerificationCode: `
			SELECT id, user_name, user_password, user_email, user_role, 
               user_approved, user_verify_code, user_verify_expires
      FROM users 
      WHERE user_verify_code = $1
    `,
	ApproveUser: `
      UPDATE users 
        SET user_approved = true, 
            user_verify_code = NULL, 
            user_verify_expires = NULL, 
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $1
        RETURNING updated_at
    `,
}
