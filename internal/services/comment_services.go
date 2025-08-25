package services

import (
	"fmt"
	"strconv"

	"goserver/internal/database"
	"goserver/internal/models"
)

func GetBlogCommentByID(commentID string) (*models.DbComment, error) {
	id, err := strconv.Atoi(commentID)
	if err != nil {
		return nil, err
	}

	var comment models.DbComment
	query := `
        SELECT id, comment_blog_id, comment_name, comment_email, comment_body, comment_approved
        FROM comments 
        WHERE id = $1
    `

	err = database.DB.Get(&comment, query, id)
	if err != nil {
		return nil, nil // Not found or decode error
	}
	return &comment, nil
}

func GetCommentsByBlogID(blogID string) ([]models.DbComment, error) {
	id, err := strconv.Atoi(blogID)
	if err != nil {
		return nil, err
	}

	var comments []models.DbComment
	query := `
        SELECT id, comment_blog_id, comment_name, comment_email, comment_body, comment_approved
        FROM comments 
        WHERE comment_blog_id = $1
        ORDER BY created_at ASC
    `

	err = database.DB.Select(&comments, query, id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// AddComment adds a new comment to the comments table
func AddComment(comment *models.DbComment) (int, error) {
	query := `
        INSERT INTO comments (comment_blog_id, comment_name, comment_email, comment_body, comment_approved) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id, created_at, updated_at
    `

	err := database.DB.QueryRowx(query,
		comment.BlogID,
		comment.Name,
		comment.Email,
		comment.Body,
		comment.Approved,
	).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		return 0, err
	}
	return comment.ID, nil
}

// UpdateComment updates a comment by its ID and blog ID
func UpdateComment(blogID, commentID string, newText string) error {
	bID, err := strconv.Atoi(blogID)
	if err != nil {
		return err
	}
	cID, err := strconv.Atoi(commentID)
	if err != nil {
		return err
	}

	// Use parameterized query to avoid SQL injection
	query := `
        UPDATE comments 
        SET comment_body = $1, updated_at = CURRENT_TIMESTAMP 
        WHERE id = $2 AND comment_blog_id = $3
    `

	_, err = database.DB.Exec(query, newText, cID, bID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteComment deletes a comment by its ID and blog ID
func DeleteComment(blogID, commentID string) error {
	bID, err := strconv.Atoi(blogID)
	if err != nil {
		return err
	}
	cID, err := strconv.Atoi(commentID)
	if err != nil {
		return err
	}

	query := `DELETE FROM comments WHERE id = $1 AND comment_blog_id = $2`

	result, err := database.DB.Exec(query, cID, bID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("comment not found")
	}
	return nil
}
