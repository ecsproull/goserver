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
	err = database.DB.Get(&comment, models.CommentQueries.GetByID, id)
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
	err = database.DB.Select(&comments, models.CommentQueries.GetByBlogID, id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func AddComment(comment *models.DbComment) (int, error) {
	err := database.DB.QueryRowx(models.CommentQueries.Insert,
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

func UpdateComment(blogID, commentID string, newText string) error {
	bID, err := strconv.Atoi(blogID)
	if err != nil {
		return err
	}
	cID, err := strconv.Atoi(commentID)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(models.CommentQueries.Update, newText, cID, bID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteComment(blogID, commentID string) error {
	bID, err := strconv.Atoi(blogID)
	if err != nil {
		return err
	}
	cID, err := strconv.Atoi(commentID)
	if err != nil {
		return err
	}

	result, err := database.DB.Exec(models.CommentQueries.Delete, cID, bID)
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
