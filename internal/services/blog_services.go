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
	err := database.DB.Select(&blogs, models.BlogQueries.GetAll)
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
	err = database.DB.Get(&blog, models.BlogQueries.GetByID, blogID)
	if err != nil {
		return nil, nil // Not found or decode error
	}
	return &blog, nil
}

// SaveBlog creates a new blog or updates an existing one based on blog_id
func SaveBlog(data *models.DbBlog) (string, error) {
	if data.ID != 0 {
		// Update existing blog
		var updatedBlog models.DbBlog
		_, err := database.DB.Exec(models.BlogQueries.Update,
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
		err := database.DB.QueryRowx(models.BlogQueries.Insert,
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

	result, err := database.DB.Exec(models.BlogQueries.Delete, blogID)
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
