package models

import (
	"time"
)

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

// CommentQueries holds SQL statements for the comments model.
type CQueries struct {
	GetAll      string
	GetByID     string
	GetByBlogID string
	Insert      string
	Update      string
	Delete      string
}

var CommentQueries = CQueries{
	GetByID: `
        SELECT id, comment_blog_id, comment_name, comment_email, comment_body, comment_approved
        FROM comments 
        WHERE id = $1
    `,
	GetByBlogID: `
        SELECT id, comment_blog_id, comment_name, comment_email, comment_body, comment_approved
        FROM comments 
        WHERE comment_blog_id = $1
        ORDER BY created_at ASC
    `,
	Insert: `
        INSERT INTO comments (comment_blog_id, comment_name, comment_email, comment_body, comment_approved) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id
    `,
	Update: `
        UPDATE comments 
        SET comment_body = $1, updated_at = CURRENT_TIMESTAMP 
        WHERE id = $2 AND comment_blog_id = $3
    `,
	Delete: `
        DELETE FROM comments
        WHERE id = $1 AND comment_blog_id = $2
    `,
}
