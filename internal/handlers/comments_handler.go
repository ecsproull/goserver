package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct{}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
}

func (h *CommentHandler) GetByBlogID(c *gin.Context) {
	blogID := c.Param("blogId")
	// TODO: Implement database fetch
	c.JSON(http.StatusOK, gin.H{
		"data":    []interface{}{},
		"message": "Comments retrieved successfully",
		"blogId":  blogID,
	})
}

func (h *CommentHandler) Create(c *gin.Context) {
	blogID := c.Param("blogId")
	// TODO: Implement comment creation
	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		"blogId":  blogID,
	})
}

func (h *CommentHandler) Update(c *gin.Context) {
	blogID := c.Param("blogId")
	id := c.Param("id")
	// TODO: Implement comment update
	c.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
		"blogId":  blogID,
		"id":      id,
	})
}

func (h *CommentHandler) Delete(c *gin.Context) {
	blogID := c.Param("blogId")
	id := c.Param("id")
	// TODO: Implement comment deletion
	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
		"blogId":  blogID,
		"id":      id,
	})
}
