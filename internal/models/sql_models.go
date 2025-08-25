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

type DbBlog struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"blog_subject" db:"blog_subject"`
	Content   string    `json:"blog_body" db:"blog_body"`
	AuthorID  string    `json:"blog_owner_name" db:"blog_owner_name"`
	Email     string    `json:"blog_owner_email" db:"blog_owner_email"`
	Category  string    `json:"blog_category" db:"blog_category"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type DbPlace struct {
	ID            int        `json:"id" db:"id"`
	PlaceName     string     `json:"place_name" db:"place_name"`
	PlaceInfo     string     `json:"place_info" db:"place_info"`
	PlaceLat      float64    `json:"place_lat" db:"place_lat"`
	PlaceLng      float64    `json:"place_lng" db:"place_lng"`
	PlaceIconType int        `json:"place_icon_type" db:"place_icon_type"`
	PlaceAddress  string     `json:"place_address" db:"place_address"`
	PlacePhone    string     `json:"place_phone" db:"place_phone"`
	PlaceEmail    string     `json:"place_email" db:"place_email"`
	PlaceWebsite  string     `json:"place_website" db:"place_website"`
	PlaceArrive   *time.Time `json:"place_arrive" db:"place_arrive"`
	PlaceDepart   *time.Time `json:"place_depart" db:"place_depart"`
	PlaceHideInfo BoolInt    `json:"place_hide_info" db:"place_hide_info"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

type DbComment struct {
	ID        int       `json:"id" db:"id"`
	BlogID    int       `json:"comment_blog_id" db:"comment_blog_id"`
	Name      string    `json:"comment_name" db:"comment_name"`
	Email     string    `json:"comment_email" db:"comment_email"`
	Body      string    `json:"comment_body" db:"comment_body"`
	Approved  bool      `json:"comment_approved" db:"comment_approved"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
