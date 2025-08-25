package services

import (
	"fmt"
	"log"
	"strconv"

	"goserver/internal/database"
	"goserver/internal/models"
)

func GetAllBlogs() ([]models.DbBlog, error) {
	var blogs []models.DbBlog

	query := `
        SELECT id, blog_subject, blog_body, blog_owner_name, blog_owner_email, blog_category
        FROM blogs
    `

	err := database.DB.Select(&blogs, query)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func GetBlogByID(id string) (*models.DbBlog, error) {
	blogID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	var blog models.DbBlog
	query := `
        SELECT id, blog_subject, blog_body, blog_owner_name, blog_owner_email, blog_category
        FROM blogs 
        WHERE id = $1
    `

	err = database.DB.Get(&blog, query, blogID)
	if err != nil {
		return nil, nil // Not found or decode error
	}
	return &blog, nil
}

// SaveBlog creates a new blog or updates an existing one based on blog_id
func SaveBlog(data *models.DbBlog) (string, error) {
	if data.ID != 0 {
		// Update existing blog
		query := `
            UPDATE blogs 
            SET blog_subject = $1, blog_body = $2, blog_owner_name = $3, blog_owner_email = $4, blog_category = $5, updated_at = CURRENT_TIMESTAMP
            WHERE id = $6
            RETURNING id, blog_subject, blog_body, blog_owner_name, blog_owner_email, blog_category
        `

		var updatedBlog models.DbBlog
		_, err := database.DB.Exec(query,
			data.Title,
			data.Content,
			data.AuthorID,
			data.Email,
			data.Category,
			data.ID,
		)

		if err != nil {
			return "", err
		}
		log.Printf("Saved Blog: %s", updatedBlog.Title)
		return strconv.Itoa(data.ID), nil
	} else {
		// Create new blog
		log.Printf("Creating new blog post.")
		query := `
            INSERT INTO blogs (blog_subject, blog_body, blog_owner_name, blog_owner_email, blog_category) 
            VALUES ($1, $2, $3, $4, $5) 
            RETURNING id
        `

		err := database.DB.QueryRowx(query,
			data.Title,
			data.Content,
			data.AuthorID,
			data.Email,
			data.Category,
		).Scan(&data.ID)

		if err != nil {
			return "", err
		}
		log.Printf("Saved Blog: %s", data.Title)
		return strconv.Itoa(data.ID), nil
	}
}

// DeleteBlog deletes a blog by its ID
func DeleteBlog(id string) error {
	blogID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM blogs WHERE id = $1`

	result, err := database.DB.Exec(query, blogID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("blog not found")
	}
	return nil
}
