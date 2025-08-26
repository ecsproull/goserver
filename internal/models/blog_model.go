package models

import (
	"time"
)

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

type BQueries struct {
	GetAll  string
	GetByID string
	Insert  string
	Update  string
	Delete  string
}

var BlogQueries = BQueries{
	GetAll: `
        SELECT id, blog_subject, blog_body, blog_owner_name, blog_owner_email, blog_category
        FROM blogs
    `,
	GetByID: `
        SELECT id, blog_subject, blog_body, blog_owner_name, blog_owner_email, blog_category
        FROM blogs 
        WHERE id = $1
    `,
	Insert: `
        INSERT INTO blogs (blog_subject, blog_body, blog_owner_name, blog_owner_email, blog_category) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id
    `,
	Update: `
        UPDATE blogs 
        SET blog_subject = $1, blog_body = $2, blog_owner_name = $3, blog_owner_email = $4, blog_category = $5, updated_at = CURRENT_TIMESTAMP
        WHERE id = $6
        RETURNING id, blog_subject, blog_body, blog_owner_name, blog_owner_email, blog_category
    `,
	Delete: `
        DELETE FROM blogs
        WHERE id = $1
    `,
}
