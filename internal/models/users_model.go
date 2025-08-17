package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"user_name" bson:"user_name"`
	Email         string             `json:"user_email,omitempty" bson:"user_email,omitempty"`
	Password      string             `json:"user_password,omitempty" bson:"user_password,omitempty"`
	Approved      bool               `json:"user_approved,omitempty" bson:"user_approved,omitempty"`
	VerifyCode    string             `json:"user_verify_code,omitempty" bson:"user_verify_code,omitempty"`
	VerifyExpires time.Time          `json:"user_verify_expires" bson:"user_verify_expires,omitempty"`
	Role          string             `json:"role" bson:"role"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}
